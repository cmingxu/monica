//------------------------------------------------------------------------------
// <auto-generated>
//     This code was generated by a tool.
//
//     Changes to this file may cause incorrect behavior and will be lost if
//     the code is regenerated.
// </auto-generated>
//------------------------------------------------------------------------------

// Generated from: protobuf_defs/gds.proto
namespace gds
{
  [global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@"Building")]
  public partial class Building : global::ProtoBuf.IExtensible
  {
    public Building() {}
    
    private string _name;
    [global::ProtoBuf.ProtoMember(1, IsRequired = true, Name=@"name", DataFormat = global::ProtoBuf.DataFormat.Default)]
    public string name
    {
      get { return _name; }
      set { _name = value; }
    }
    private uint _level;
    [global::ProtoBuf.ProtoMember(2, IsRequired = true, Name=@"level", DataFormat = global::ProtoBuf.DataFormat.TwosComplement)]
    public uint level
    {
      get { return _level; }
      set { _level = value; }
    }
    private string _icon;
    [global::ProtoBuf.ProtoMember(3, IsRequired = true, Name=@"icon", DataFormat = global::ProtoBuf.DataFormat.Default)]
    public string icon
    {
      get { return _icon; }
      set { _icon = value; }
    }
    private string _model;
    [global::ProtoBuf.ProtoMember(4, IsRequired = true, Name=@"model", DataFormat = global::ProtoBuf.DataFormat.Default)]
    public string model
    {
      get { return _model; }
      set { _model = value; }
    }
    private string _texture;
    [global::ProtoBuf.ProtoMember(5, IsRequired = true, Name=@"texture", DataFormat = global::ProtoBuf.DataFormat.Default)]
    public string texture
    {
      get { return _texture; }
      set { _texture = value; }
    }
    private global::ProtoBuf.IExtension extensionObject;
    global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
      { return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }
  }
  
  [global::System.Serializable, global::ProtoBuf.ProtoContract(Name=@"BuildingGds")]
  public partial class BuildingGds : global::ProtoBuf.IExtensible
  {
    public BuildingGds() {}
    
    private readonly global::System.Collections.Generic.List<gds.Building> _buildings = new global::System.Collections.Generic.List<gds.Building>();
    [global::ProtoBuf.ProtoMember(1, Name=@"buildings", DataFormat = global::ProtoBuf.DataFormat.Default)]
    public global::System.Collections.Generic.List<gds.Building> buildings
    {
      get { return _buildings; }
    }
  
    private global::ProtoBuf.IExtension extensionObject;
    global::ProtoBuf.IExtension global::ProtoBuf.IExtensible.GetExtensionObject(bool createIfMissing)
      { return global::ProtoBuf.Extensible.GetExtensionObject(ref extensionObject, createIfMissing); }
  }
  
}