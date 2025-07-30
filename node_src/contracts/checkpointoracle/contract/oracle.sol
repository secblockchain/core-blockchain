pragma solidity ^0.6.0;

/**
 * @title CheckpointOracle
 * @author Gary Rong<garyrong@ethereum.org>, Martin Swende <martin.swende@ethereum.org>
 * @dev Implementation of the blockchain checkpoint registrar.
 */
contract CheckpointOracle {
    /*
        Events
    */

    // NewCheckpointVote is emitted when a new checkpoint proposal receives a vote.
    event NewCheckpointVote(
        uint64 indexed index,
        bytes32 checkpointHash,
        uint8 v,
        bytes32 r,
        bytes32 s
    );

    /*
        Fields
    */
    // A map of admin users who have the permission to update CHT and bloom Trie root
    mapping(address => bool) admins;

    // A list of admin users so that we can obtain all admin users.
    address[] adminList;

    // Latest stored section id
    uint64 sectionIndex;

    // The block height associated with latest registered checkpoint.
    uint height;

    // The hash of latest registered checkpoint.
    bytes32 hash;

    // The frequency for creating a checkpoint
    //
    // The default value should be the same as the checkpoint size(32768) in the ethereum.
    uint sectionSize;

    // The number of confirmations needed before a checkpoint can be registered.
    // We have to make sure the checkpoint registered will not be invalid due to
    // chain reorg.
    //
    // The default value should be the same as the checkpoint process confirmations(256)
    // in the ethereum.
    uint processConfirms;

    // The required signatures to finalize a stable checkpoint.
    uint threshold;

    /*
        EIP-712 Domain Separator
    */
    bytes32 public constant DOMAIN_SEPARATOR_TYPEHASH =
        keccak256(
            "EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"
        );
    bytes32 public constant CHECKPOINT_TYPEHASH =
        keccak256("Checkpoint(uint64 sectionIndex,bytes32 hash,uint256 nonce)");

    string public constant DOMAIN_NAME = "CheckpointOracle";
    string public constant DOMAIN_VERSION = "1.0";
    uint256 public immutable chainId;

    // Per-admin nonces for replay protection
    mapping(address => uint256) public adminNonces;

    /*
        Public Functions
    */
    constructor(
        address[] memory _adminlist,
        uint _sectionSize,
        uint _processConfirms,
        uint _threshold
    ) public {
        for (uint i = 0; i < _adminlist.length; i++) {
            admins[_adminlist[i]] = true;
            adminList.push(_adminlist[i]);
        }
        sectionSize = _sectionSize;
        processConfirms = _processConfirms;
        threshold = _threshold;

        // Get chain ID for EIP-712 domain separator
        uint256 id;
        assembly {
            id := chainid()
        }
        chainId = id;
    }

    /**
     * @dev Get latest stable checkpoint information.
     * @return section index
     * @return checkpoint hash
     * @return block height associated with checkpoint
     */
    function GetLatestCheckpoint() public view returns (uint64, bytes32, uint) {
        return (sectionIndex, hash, height);
    }

    /**
     * @dev Get the current nonce for an admin address
     * @param admin The admin address
     * @return The current nonce
     */
    function getAdminNonce(address admin) public view returns (uint256) {
        return adminNonces[admin];
    }

    // SetCheckpoint sets a new checkpoint using EIP-712 signatures
    // @_hash : the hash to set at _sectionIndex
    // @_sectionIndex : the section index to set
    // @_nonces : the nonces used by each signer for replay protection
    // @v : the list of v-values
    // @r : the list or r-values
    // @s : the list of s-values
    function SetCheckpoint(
        bytes32 _hash,
        uint64 _sectionIndex,
        uint256[] memory _nonces,
        uint8[] memory v,
        bytes32[] memory r,
        bytes32[] memory s
    ) public returns (bool) {
        // Ensure the sender is authorized.
        require(admins[msg.sender]);

        // Ensure the batch of signatures are valid.
        require(v.length == r.length);
        require(v.length == s.length);
        require(v.length == _nonces.length);

        // Filter out "future" checkpoint.
        if (
            block.number < (_sectionIndex + 1) * sectionSize + processConfirms
        ) {
            return false;
        }
        // Filter out "old" announcement
        if (_sectionIndex < sectionIndex) {
            return false;
        }
        // Filter out "stale" announcement
        if (
            _sectionIndex == sectionIndex && (_sectionIndex != 0 || height != 0)
        ) {
            return false;
        }
        // Filter out "invalid" announcement
        if (_hash == "") {
            return false;
        }

        // EIP-712 domain separator
        bytes32 domainSeparator = keccak256(abi.encode(
            DOMAIN_SEPARATOR_TYPEHASH,
            keccak256(bytes(DOMAIN_NAME)),
            keccak256(bytes(DOMAIN_VERSION)),
            chainId,
            address(this)
        ));

        address lastVoter = address(0);

        // In order for us not to have to maintain a mapping of who has already
        // voted, and we don't want to count a vote twice, the signatures must
        // be submitted in strict ordering.
        for (uint idx = 0; idx < v.length; idx++) {
            // Create the checkpoint struct hash
            bytes32 checkpointHash = keccak256(abi.encode(
                CHECKPOINT_TYPEHASH,
                _sectionIndex,
                _hash,
                _nonces[idx]
            ));
            
            // Create the final hash to sign (EIP-712 format)
            bytes32 signedHash = keccak256(abi.encodePacked(
                "\x19\x01",
                domainSeparator,
                checkpointHash
            ));

            address signer = ecrecover(signedHash, v[idx], r[idx], s[idx]);
            require(admins[signer]);
            require(uint256(signer) > uint256(lastVoter));

            // Verify nonce and increment it
            require(adminNonces[signer] == _nonces[idx], "Invalid nonce");
            adminNonces[signer]++;

            lastVoter = signer;
            emit NewCheckpointVote(
                _sectionIndex,
                _hash,
                v[idx],
                r[idx],
                s[idx]
            );

            // Sufficient signatures present, update latest checkpoint.
            if (idx + 1 >= threshold) {
                hash = _hash;
                height = block.number;
                sectionIndex = _sectionIndex;
                return true;
            }
        }
        // We shouldn't wind up here, reverting un-emits the events
        revert();
    }

    /**
     * @dev Get all admin addresses
     * @return address list
     */
    function GetAllAdmin() public view returns (address[] memory) {
        address[] memory ret = new address[](adminList.length);
        for (uint i = 0; i < adminList.length; i++) {
            ret[i] = adminList[i];
        }
        return ret;
    }
}
