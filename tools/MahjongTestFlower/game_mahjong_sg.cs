//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

// Generated from: game_mahjong_sg.proto
namespace mahjong
{
    [global::ProtoBuf.ProtoContract(Name=@"SGGreatWinType")]
    public enum SGGreatWinType
    {
            
      [global::ProtoBuf.ProtoEnum(Name=@"SGGreatWinType_None", Value=0)]
      SGGreatWinType_None = 0,
            
      [global::ProtoBuf.ProtoEnum(Name=@"SGGreatWinType_PongKong", Value=1)]
      SGGreatWinType_PongKong = 1,
            
      [global::ProtoBuf.ProtoEnum(Name=@"SGGreatWinType_PureSame", Value=2)]
      SGGreatWinType_PureSame = 2,
            
      [global::ProtoBuf.ProtoEnum(Name=@"SGGreatWinType_MixedSame", Value=4)]
      SGGreatWinType_MixedSame = 4,
            
      [global::ProtoBuf.ProtoEnum(Name=@"SGGreatWinType_MixPongKong", Value=8)]
      SGGreatWinType_MixPongKong = 8,
            
      [global::ProtoBuf.ProtoEnum(Name=@"SGGreatWinType_PurePongKong", Value=32)]
      SGGreatWinType_PurePongKong = 32,
            
      [global::ProtoBuf.ProtoEnum(Name=@"SGGreatWinType_Suit19", Value=64)]
      SGGreatWinType_Suit19 = 64,
            
      [global::ProtoBuf.ProtoEnum(Name=@"SGGreatWinType_AllHonor", Value=128)]
      SGGreatWinType_AllHonor = 128,
            
      [global::ProtoBuf.ProtoEnum(Name=@"SGGreatWinType_ThirteenOrphans", Value=256)]
      SGGreatWinType_ThirteenOrphans = 256
    }
  
    [global::ProtoBuf.ProtoContract(Name=@"SGHandOverExtraFlags")]
    public enum SGHandOverExtraFlags
    {
            
      [global::ProtoBuf.ProtoEnum(Name=@"SGHOFlags_None", Value=0)]
      SGHOFlags_None = 0,
            
      [global::ProtoBuf.ProtoEnum(Name=@"SGHOFlags_Faker", Value=1)]
      SGHOFlags_Faker = 1,
            
      [global::ProtoBuf.ProtoEnum(Name=@"SGHOFlags_FollowBanker", Value=2)]
      SGHOFlags_FollowBanker = 2,
            
      [global::ProtoBuf.ProtoEnum(Name=@"SGHOFlags_ExposedKong2SelfDrawn", Value=4)]
      SGHOFlags_ExposedKong2SelfDrawn = 4,
            
      [global::ProtoBuf.ProtoEnum(Name=@"SGHOFlags_SelfDrawn", Value=8)]
      SGHOFlags_SelfDrawn = 8
    }
  
}