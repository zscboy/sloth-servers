//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

// Generated from: game_pokerface_split2.proto
namespace pokerface
{
  [global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@"MsgPlayerInfo")]
  public partial class MsgPlayerInfo : global::ProtoBuf.IExtensible
  {
    public MsgPlayerInfo() {}
    
    private string _userID;
    [global::ProtoBuf.ProtoMember(1, IsRequired = true, Name=@"userID", DataFormat = global::ProtoBuf.DataFormat.Default)]
    public string userID
    {
      get { return _userID; }
      set { _userID = value; }
    }
    private int _chairID;
    [global::ProtoBuf.ProtoMember(2, IsRequired = true, Name=@"chairID", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    public int chairID
    {
      get { return _chairID; }
      set { _chairID = value; }
    }
    private int _state;
    [global::ProtoBuf.ProtoMember(3, IsRequired = true, Name=@"state", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    public int state
    {
      get { return _state; }
      set { _state = value; }
    }

    private string _name = "";
    [global::ProtoBuf.ProtoMember(4, IsRequired = false, Name=@"name", DataFormat = global::ProtoBuf.DataFormat.Default)]
    [global::System.ComponentModel.DefaultValue("")]
    public string name
    {
      get { return _name; }
      set { _name = value; }
    }

    private string _nick = "";
    [global::ProtoBuf.ProtoMember(5, IsRequired = false, Name=@"nick", DataFormat = global::ProtoBuf.DataFormat.Default)]
    [global::System.ComponentModel.DefaultValue("")]
    public string nick
    {
      get { return _nick; }
      set { _nick = value; }
    }

    private uint _sex = default(uint);
    [global::ProtoBuf.ProtoMember(6, IsRequired = false, Name=@"sex", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    [global::System.ComponentModel.DefaultValue(default(uint))]
    public uint sex
    {
      get { return _sex; }
      set { _sex = value; }
    }

    private string _headIconURI = "";
    [global::ProtoBuf.ProtoMember(7, IsRequired = false, Name=@"headIconURI", DataFormat = global::ProtoBuf.DataFormat.Default)]
    [global::System.ComponentModel.DefaultValue("")]
    public string headIconURI
    {
      get { return _headIconURI; }
      set { _headIconURI = value; }
    }

    private string _ip = "";
    [global::ProtoBuf.ProtoMember(8, IsRequired = false, Name=@"ip", DataFormat = global::ProtoBuf.DataFormat.Default)]
    [global::System.ComponentModel.DefaultValue("")]
    public string ip
    {
      get { return _ip; }
      set { _ip = value; }
    }

    private string _location = "";
    [global::ProtoBuf.ProtoMember(9, IsRequired = false, Name=@"location", DataFormat = global::ProtoBuf.DataFormat.Default)]
    [global::System.ComponentModel.DefaultValue("")]
    public string location
    {
      get { return _location; }
      set { _location = value; }
    }

    private int _dfHands = default(int);
    [global::ProtoBuf.ProtoMember(10, IsRequired = false, Name=@"dfHands", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    [global::System.ComponentModel.DefaultValue(default(int))]
    public int dfHands
    {
      get { return _dfHands; }
      set { _dfHands = value; }
    }

    private int _diamond = default(int);
    [global::ProtoBuf.ProtoMember(11, IsRequired = false, Name=@"diamond", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    [global::System.ComponentModel.DefaultValue(default(int))]
    public int diamond
    {
      get { return _diamond; }
      set { _diamond = value; }
    }

    private int _charm = default(int);
    [global::ProtoBuf.ProtoMember(12, IsRequired = false, Name=@"charm", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    [global::System.ComponentModel.DefaultValue(default(int))]
    public int charm
    {
      get { return _charm; }
      set { _charm = value; }
    }

    private int _avatarID = default(int);
    [global::ProtoBuf.ProtoMember(13, IsRequired = false, Name=@"avatarID", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    [global::System.ComponentModel.DefaultValue(default(int))]
    public int avatarID
    {
      get { return _avatarID; }
      set { _avatarID = value; }
    }
    private global::ProtoBuf.IExtension extensionObject;
    global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
      { return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }
  }
  
  [global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@"PlayerHandScoreRecord")]
  public partial class PlayerHandScoreRecord : global::ProtoBuf.IExtensible
  {
    public PlayerHandScoreRecord() {}
    
    private string _userID;
    [global::ProtoBuf.ProtoMember(1, IsRequired = true, Name=@"userID", DataFormat = global::ProtoBuf.DataFormat.Default)]
    public string userID
    {
      get { return _userID; }
      set { _userID = value; }
    }
    private int _winType;
    [global::ProtoBuf.ProtoMember(2, IsRequired = true, Name=@"winType", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    public int winType
    {
      get { return _winType; }
      set { _winType = value; }
    }
    private int _score;
    [global::ProtoBuf.ProtoMember(3, IsRequired = true, Name=@"score", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    public int score
    {
      get { return _score; }
      set { _score = value; }
    }
    private global::ProtoBuf.IExtension extensionObject;
    global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
      { return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }
  }
  
  [global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@"MsgRoomHandScoreRecord")]
  public partial class MsgRoomHandScoreRecord : global::ProtoBuf.IExtensible
  {
    public MsgRoomHandScoreRecord() {}
    
    private int _endType;
    [global::ProtoBuf.ProtoMember(1, IsRequired = true, Name=@"endType", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    public int endType
    {
      get { return _endType; }
      set { _endType = value; }
    }
    private int _handIndex;
    [global::ProtoBuf.ProtoMember(2, IsRequired = true, Name=@"handIndex", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    public int handIndex
    {
      get { return _handIndex; }
      set { _handIndex = value; }
    }
    private readonly global::System.Collections.Generic.List<pokerface.PlayerHandScoreRecord> _playerRecords = new global::System.Collections.Generic.List<pokerface.PlayerHandScoreRecord>();
    [global::ProtoBuf.ProtoMember(3, Name=@"playerRecords", DataFormat = global::ProtoBuf.DataFormat.Default)]
    public global::System.Collections.Generic.List<pokerface.PlayerHandScoreRecord> playerRecords
    {
      get { return _playerRecords; }
    }
  
    private global::ProtoBuf.IExtension extensionObject;
    global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
      { return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }
  }
  
  [global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@"MsgRoomInfo")]
  public partial class MsgRoomInfo : global::ProtoBuf.IExtensible
  {
    public MsgRoomInfo() {}
    
    private int _state;
    [global::ProtoBuf.ProtoMember(1, IsRequired = true, Name=@"state", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    public int state
    {
      get { return _state; }
      set { _state = value; }
    }
    private readonly global::System.Collections.Generic.List<pokerface.MsgPlayerInfo> _players = new global::System.Collections.Generic.List<pokerface.MsgPlayerInfo>();
    [global::ProtoBuf.ProtoMember(2, Name=@"players", DataFormat = global::ProtoBuf.DataFormat.Default)]
    public global::System.Collections.Generic.List<pokerface.MsgPlayerInfo> players
    {
      get { return _players; }
    }
  

    private string _ownerID = "";
    [global::ProtoBuf.ProtoMember(3, IsRequired = false, Name=@"ownerID", DataFormat = global::ProtoBuf.DataFormat.Default)]
    [global::System.ComponentModel.DefaultValue("")]
    public string ownerID
    {
      get { return _ownerID; }
      set { _ownerID = value; }
    }

    private string _roomNumber = "";
    [global::ProtoBuf.ProtoMember(4, IsRequired = false, Name=@"roomNumber", DataFormat = global::ProtoBuf.DataFormat.Default)]
    [global::System.ComponentModel.DefaultValue("")]
    public string roomNumber
    {
      get { return _roomNumber; }
      set { _roomNumber = value; }
    }

    private int _handStartted = default(int);
    [global::ProtoBuf.ProtoMember(5, IsRequired = false, Name=@"handStartted", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    [global::System.ComponentModel.DefaultValue(default(int))]
    public int handStartted
    {
      get { return _handStartted; }
      set { _handStartted = value; }
    }
    private readonly global::System.Collections.Generic.List<pokerface.MsgRoomHandScoreRecord> _scoreRecords = new global::System.Collections.Generic.List<pokerface.MsgRoomHandScoreRecord>();
    [global::ProtoBuf.ProtoMember(6, Name=@"scoreRecords", DataFormat = global::ProtoBuf.DataFormat.Default)]
    public global::System.Collections.Generic.List<pokerface.MsgRoomHandScoreRecord> scoreRecords
    {
      get { return _scoreRecords; }
    }
  
    private global::ProtoBuf.IExtension extensionObject;
    global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
      { return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }
  }
  
  [global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@"RoomScoreRecords")]
  public partial class RoomScoreRecords : global::ProtoBuf.IExtensible
  {
    public RoomScoreRecords() {}
    
    private readonly global::System.Collections.Generic.List<pokerface.MsgRoomHandScoreRecord> _scoreRecords = new global::System.Collections.Generic.List<pokerface.MsgRoomHandScoreRecord>();
    [global::ProtoBuf.ProtoMember(1, Name=@"scoreRecords", DataFormat = global::ProtoBuf.DataFormat.Default)]
    public global::System.Collections.Generic.List<pokerface.MsgRoomHandScoreRecord> scoreRecords
    {
      get { return _scoreRecords; }
    }
  
    private global::ProtoBuf.IExtension extensionObject;
    global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
      { return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }
  }
  
  [global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@"MsgDisbandAnswer")]
  public partial class MsgDisbandAnswer : global::ProtoBuf.IExtensible
  {
    public MsgDisbandAnswer() {}
    
    private bool _agree;
    [global::ProtoBuf.ProtoMember(1, IsRequired = true, Name=@"agree", DataFormat = global::ProtoBuf.DataFormat.Default)]
    public bool agree
    {
      get { return _agree; }
      set { _agree = value; }
    }
    private global::ProtoBuf.IExtension extensionObject;
    global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
      { return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }
  }
  
  [global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@"MsgDisbandNotify")]
  public partial class MsgDisbandNotify : global::ProtoBuf.IExtensible
  {
    public MsgDisbandNotify() {}
    
    private int _disbandState;
    [global::ProtoBuf.ProtoMember(1, IsRequired = true, Name=@"disbandState", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    public int disbandState
    {
      get { return _disbandState; }
      set { _disbandState = value; }
    }
    private int _applicant;
    [global::ProtoBuf.ProtoMember(2, IsRequired = true, Name=@"applicant", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    public int applicant
    {
      get { return _applicant; }
      set { _applicant = value; }
    }
    private readonly global::System.Collections.Generic.List<int> _waits = new global::System.Collections.Generic.List<int>();
    [global::ProtoBuf.ProtoMember(3, Name=@"waits", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    public global::System.Collections.Generic.List<int> waits
    {
      get { return _waits; }
    }
  
    private readonly global::System.Collections.Generic.List<int> _agrees = new global::System.Collections.Generic.List<int>();
    [global::ProtoBuf.ProtoMember(4, Name=@"agrees", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    public global::System.Collections.Generic.List<int> agrees
    {
      get { return _agrees; }
    }
  
    private readonly global::System.Collections.Generic.List<int> _rejects = new global::System.Collections.Generic.List<int>();
    [global::ProtoBuf.ProtoMember(5, Name=@"rejects", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    public global::System.Collections.Generic.List<int> rejects
    {
      get { return _rejects; }
    }
  

    private int _countdown = default(int);
    [global::ProtoBuf.ProtoMember(6, IsRequired = false, Name=@"countdown", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    [global::System.ComponentModel.DefaultValue(default(int))]
    public int countdown
    {
      get { return _countdown; }
      set { _countdown = value; }
    }
    private global::ProtoBuf.IExtension extensionObject;
    global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
      { return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }
  }
  
  [global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@"MsgGameOverPlayerStat")]
  public partial class MsgGameOverPlayerStat : global::ProtoBuf.IExtensible
  {
    public MsgGameOverPlayerStat() {}
    
    private int _chairID;
    [global::ProtoBuf.ProtoMember(1, IsRequired = true, Name=@"chairID", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    public int chairID
    {
      get { return _chairID; }
      set { _chairID = value; }
    }
    private int _score;
    [global::ProtoBuf.ProtoMember(2, IsRequired = true, Name=@"score", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    public int score
    {
      get { return _score; }
      set { _score = value; }
    }
    private int _winChuckCounter;
    [global::ProtoBuf.ProtoMember(3, IsRequired = true, Name=@"winChuckCounter", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    public int winChuckCounter
    {
      get { return _winChuckCounter; }
      set { _winChuckCounter = value; }
    }
    private int _winSelfDrawnCounter;
    [global::ProtoBuf.ProtoMember(4, IsRequired = true, Name=@"winSelfDrawnCounter", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    public int winSelfDrawnCounter
    {
      get { return _winSelfDrawnCounter; }
      set { _winSelfDrawnCounter = value; }
    }
    private int _chuckerCounter;
    [global::ProtoBuf.ProtoMember(5, IsRequired = true, Name=@"chuckerCounter", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    public int chuckerCounter
    {
      get { return _chuckerCounter; }
      set { _chuckerCounter = value; }
    }

    private int _robKongCounter = default(int);
    [global::ProtoBuf.ProtoMember(6, IsRequired = false, Name=@"robKongCounter", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    [global::System.ComponentModel.DefaultValue(default(int))]
    public int robKongCounter
    {
      get { return _robKongCounter; }
      set { _robKongCounter = value; }
    }

    private int _kongerCounter = default(int);
    [global::ProtoBuf.ProtoMember(7, IsRequired = false, Name=@"kongerCounter", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    [global::System.ComponentModel.DefaultValue(default(int))]
    public int kongerCounter
    {
      get { return _kongerCounter; }
      set { _kongerCounter = value; }
    }
    private global::ProtoBuf.IExtension extensionObject;
    global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
      { return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }
  }
  
  [global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@"MsgGameOver")]
  public partial class MsgGameOver : global::ProtoBuf.IExtensible
  {
    public MsgGameOver() {}
    
    private readonly global::System.Collections.Generic.List<pokerface.MsgGameOverPlayerStat> _playerStats = new global::System.Collections.Generic.List<pokerface.MsgGameOverPlayerStat>();
    [global::ProtoBuf.ProtoMember(1, Name=@"playerStats", DataFormat = global::ProtoBuf.DataFormat.Default)]
    public global::System.Collections.Generic.List<pokerface.MsgGameOverPlayerStat> playerStats
    {
      get { return _playerStats; }
    }
  
    private global::ProtoBuf.IExtension extensionObject;
    global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
      { return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }
  }
  
  [global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@"MsgRoomShowTips")]
  public partial class MsgRoomShowTips : global::ProtoBuf.IExtensible
  {
    public MsgRoomShowTips() {}
    

    private string _tips = "";
    [global::ProtoBuf.ProtoMember(1, IsRequired = false, Name=@"tips", DataFormat = global::ProtoBuf.DataFormat.Default)]
    [global::System.ComponentModel.DefaultValue("")]
    public string tips
    {
      get { return _tips; }
      set { _tips = value; }
    }
    private int _tipCode;
    [global::ProtoBuf.ProtoMember(2, IsRequired = true, Name=@"tipCode", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    public int tipCode
    {
      get { return _tipCode; }
      set { _tipCode = value; }
    }
    private global::ProtoBuf.IExtension extensionObject;
    global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
      { return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }
  }
  
  [global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@"MsgRoomDelete")]
  public partial class MsgRoomDelete : global::ProtoBuf.IExtensible
  {
    public MsgRoomDelete() {}
    
    private int _reason;
    [global::ProtoBuf.ProtoMember(1, IsRequired = true, Name=@"reason", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    public int reason
    {
      get { return _reason; }
      set { _reason = value; }
    }
    private global::ProtoBuf.IExtension extensionObject;
    global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
      { return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }
  }
  
  [global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@"MsgKickout")]
  public partial class MsgKickout : global::ProtoBuf.IExtensible
  {
    public MsgKickout() {}
    
    private string _victimUserID;
    [global::ProtoBuf.ProtoMember(1, IsRequired = true, Name=@"victimUserID", DataFormat = global::ProtoBuf.DataFormat.Default)]
    public string victimUserID
    {
      get { return _victimUserID; }
      set { _victimUserID = value; }
    }
    private global::ProtoBuf.IExtension extensionObject;
    global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
      { return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }
  }
  
  [global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@"MsgKickoutResult")]
  public partial class MsgKickoutResult : global::ProtoBuf.IExtensible
  {
    public MsgKickoutResult() {}
    
    private int _result;
    [global::ProtoBuf.ProtoMember(1, IsRequired = true, Name=@"result", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    public int result
    {
      get { return _result; }
      set { _result = value; }
    }

    private string _victimUserID = "";
    [global::ProtoBuf.ProtoMember(2, IsRequired = false, Name=@"victimUserID", DataFormat = global::ProtoBuf.DataFormat.Default)]
    [global::System.ComponentModel.DefaultValue("")]
    public string victimUserID
    {
      get { return _victimUserID; }
      set { _victimUserID = value; }
    }

    private string _victimNick = "";
    [global::ProtoBuf.ProtoMember(3, IsRequired = false, Name=@"victimNick", DataFormat = global::ProtoBuf.DataFormat.Default)]
    [global::System.ComponentModel.DefaultValue("")]
    public string victimNick
    {
      get { return _victimNick; }
      set { _victimNick = value; }
    }

    private string _byWhoNick = "";
    [global::ProtoBuf.ProtoMember(4, IsRequired = false, Name=@"byWhoNick", DataFormat = global::ProtoBuf.DataFormat.Default)]
    [global::System.ComponentModel.DefaultValue("")]
    public string byWhoNick
    {
      get { return _byWhoNick; }
      set { _byWhoNick = value; }
    }

    private string _byWhoUserID = "";
    [global::ProtoBuf.ProtoMember(5, IsRequired = false, Name=@"byWhoUserID", DataFormat = global::ProtoBuf.DataFormat.Default)]
    [global::System.ComponentModel.DefaultValue("")]
    public string byWhoUserID
    {
      get { return _byWhoUserID; }
      set { _byWhoUserID = value; }
    }
    private global::ProtoBuf.IExtension extensionObject;
    global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
      { return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }
  }
  
  [global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@"MsgEnterRoomResult")]
  public partial class MsgEnterRoomResult : global::ProtoBuf.IExtensible
  {
    public MsgEnterRoomResult() {}
    
    private int _status;
    [global::ProtoBuf.ProtoMember(1, IsRequired = true, Name=@"status", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    public int status
    {
      get { return _status; }
      set { _status = value; }
    }
    private global::ProtoBuf.IExtension extensionObject;
    global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
      { return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }
  }
  
  [global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@"MsgDonate")]
  public partial class MsgDonate : global::ProtoBuf.IExtensible
  {
    public MsgDonate() {}
    
    private int _toChairID;
    [global::ProtoBuf.ProtoMember(1, IsRequired = true, Name=@"toChairID", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    public int toChairID
    {
      get { return _toChairID; }
      set { _toChairID = value; }
    }
    private int _itemID;
    [global::ProtoBuf.ProtoMember(2, IsRequired = true, Name=@"itemID", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    public int itemID
    {
      get { return _itemID; }
      set { _itemID = value; }
    }

    private int _fromChairID = default(int);
    [global::ProtoBuf.ProtoMember(3, IsRequired = false, Name=@"fromChairID", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    [global::System.ComponentModel.DefaultValue(default(int))]
    public int fromChairID
    {
      get { return _fromChairID; }
      set { _fromChairID = value; }
    }
    private global::ProtoBuf.IExtension extensionObject;
    global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
      { return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }
  }
  
    [global::ProtoBuf.ProtoContract(Name=@"RoomState")]
    public enum RoomState
    {
            
      [global::ProtoBuf.ProtoEnum(Name=@"SRoomIdle", Value=0)]
      SRoomIdle = 0,
            
      [global::ProtoBuf.ProtoEnum(Name=@"SRoomWaiting", Value=1)]
      SRoomWaiting = 1,
            
      [global::ProtoBuf.ProtoEnum(Name=@"SRoomPlaying", Value=2)]
      SRoomPlaying = 2,
            
      [global::ProtoBuf.ProtoEnum(Name=@"SRoomDeleted", Value=3)]
      SRoomDeleted = 3
    }
  
    [global::ProtoBuf.ProtoContract(Name=@"PlayerState")]
    public enum PlayerState
    {
            
      [global::ProtoBuf.ProtoEnum(Name=@"PSNone", Value=0)]
      PSNone = 0,
            
      [global::ProtoBuf.ProtoEnum(Name=@"PSReady", Value=1)]
      PSReady = 1,
            
      [global::ProtoBuf.ProtoEnum(Name=@"PSOffline", Value=2)]
      PSOffline = 2,
            
      [global::ProtoBuf.ProtoEnum(Name=@"PSPlaying", Value=3)]
      PSPlaying = 3
    }
  
    [global::ProtoBuf.ProtoContract(Name=@"DisbandState")]
    public enum DisbandState
    {
            
      [global::ProtoBuf.ProtoEnum(Name=@"Waiting", Value=1)]
      Waiting = 1,
            
      [global::ProtoBuf.ProtoEnum(Name=@"Done", Value=2)]
      Done = 2,
            
      [global::ProtoBuf.ProtoEnum(Name=@"DoneWithOtherReject", Value=3)]
      DoneWithOtherReject = 3,
            
      [global::ProtoBuf.ProtoEnum(Name=@"DoneWithRoomServerNotResponse", Value=4)]
      DoneWithRoomServerNotResponse = 4,
            
      [global::ProtoBuf.ProtoEnum(Name=@"DoneWithWaitReplyTimeout", Value=5)]
      DoneWithWaitReplyTimeout = 5,
            
      [global::ProtoBuf.ProtoEnum(Name=@"ErrorDuplicateAcquire", Value=6)]
      ErrorDuplicateAcquire = 6,
            
      [global::ProtoBuf.ProtoEnum(Name=@"ErrorNeedOwnerWhenGameNotStart", Value=7)]
      ErrorNeedOwnerWhenGameNotStart = 7
    }
  
    [global::ProtoBuf.ProtoContract(Name=@"TipCode")]
    public enum TipCode
    {
            
      [global::ProtoBuf.ProtoEnum(Name=@"TCNone", Value=0)]
      TCNone = 0,
            
      [global::ProtoBuf.ProtoEnum(Name=@"TCWaitOpponentsAction", Value=1)]
      TCWaitOpponentsAction = 1,
            
      [global::ProtoBuf.ProtoEnum(Name=@"TCDonateFailedNoEnoughDiamond", Value=2)]
      TCDonateFailedNoEnoughDiamond = 2
    }
  
    [global::ProtoBuf.ProtoContract(Name=@"RoomDeleteReason")]
    public enum RoomDeleteReason
    {
            
      [global::ProtoBuf.ProtoEnum(Name=@"IdleTimeout", Value=1)]
      IdleTimeout = 1,
            
      [global::ProtoBuf.ProtoEnum(Name=@"DisbandByOwnerFromRMS", Value=2)]
      DisbandByOwnerFromRMS = 2,
            
      [global::ProtoBuf.ProtoEnum(Name=@"DisbandByApplication", Value=3)]
      DisbandByApplication = 3,
            
      [global::ProtoBuf.ProtoEnum(Name=@"DisbandBySystem", Value=4)]
      DisbandBySystem = 4,
            
      [global::ProtoBuf.ProtoEnum(Name=@"DisbandMaxHand", Value=5)]
      DisbandMaxHand = 5,
            
      [global::ProtoBuf.ProtoEnum(Name=@"DisbandInLoseProtected", Value=6)]
      DisbandInLoseProtected = 6
    }
  
    [global::ProtoBuf.ProtoContract(Name=@"KickoutResult")]
    public enum KickoutResult
    {
            
      [global::ProtoBuf.ProtoEnum(Name=@"KickoutResult_Success", Value=1)]
      KickoutResult_Success = 1,
            
      [global::ProtoBuf.ProtoEnum(Name=@"KickoutResult_FailedGameHasStartted", Value=2)]
      KickoutResult_FailedGameHasStartted = 2,
            
      [global::ProtoBuf.ProtoEnum(Name=@"KickoutResult_FailedNeedOwner", Value=3)]
      KickoutResult_FailedNeedOwner = 3,
            
      [global::ProtoBuf.ProtoEnum(Name=@"KickoutResult_FailedPlayerNotExist", Value=4)]
      KickoutResult_FailedPlayerNotExist = 4
    }
  
    [global::ProtoBuf.ProtoContract(Name=@"EnterRoomStatus")]
    public enum EnterRoomStatus
    {
            
      [global::ProtoBuf.ProtoEnum(Name=@"Success", Value=0)]
      Success = 0,
            
      [global::ProtoBuf.ProtoEnum(Name=@"RoomNotExist", Value=1)]
      RoomNotExist = 1,
            
      [global::ProtoBuf.ProtoEnum(Name=@"RoomIsFulled", Value=2)]
      RoomIsFulled = 2,
            
      [global::ProtoBuf.ProtoEnum(Name=@"RoomPlaying", Value=3)]
      RoomPlaying = 3,
            
      [global::ProtoBuf.ProtoEnum(Name=@"InAnotherRoom", Value=4)]
      InAnotherRoom = 4,
            
      [global::ProtoBuf.ProtoEnum(Name=@"MonkeyRoomUserIDNotMatch", Value=5)]
      MonkeyRoomUserIDNotMatch = 5,
            
      [global::ProtoBuf.ProtoEnum(Name=@"MonkeyRoomUserLoginSeqNotMatch", Value=6)]
      MonkeyRoomUserLoginSeqNotMatch = 6,
            
      [global::ProtoBuf.ProtoEnum(Name=@"AppModuleNeedUpgrade", Value=7)]
      AppModuleNeedUpgrade = 7,
            
      [global::ProtoBuf.ProtoEnum(Name=@"InRoomBlackList", Value=8)]
      InRoomBlackList = 8,
            
      [global::ProtoBuf.ProtoEnum(Name=@"TakeoffDiamondFailedNotEnough", Value=9)]
      TakeoffDiamondFailedNotEnough = 9,
            
      [global::ProtoBuf.ProtoEnum(Name=@"TakeoffDiamondFailedIO", Value=10)]
      TakeoffDiamondFailedIO = 10
    }
  
}