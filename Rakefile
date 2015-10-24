task :gen do
  proto_file_path = "proto"
  proto_gen_go_path = "monica"

  Dir[proto_file_path + "/**/*.proto"].each do |p|
    dest_dir = "#{proto_gen_go_path}/#{File.dirname(p).sub('./proto', '')}"
    FileUtils.mkdir_p dest_dir

    `protoc --go_out=#{proto_gen_go_path}/#{File.dirname(p).sub("./proto", "")} --proto_path=proto/common  --proto_path=proto #{p}`
  end
end

