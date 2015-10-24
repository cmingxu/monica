task :gen do
  proto_file_path = "proto"
  proto_gen_go_path = "monica"

  Dir[proto_file_path + "/*.proto"].each do |p|
    dest_dir = "#{proto_gen_go_path}/proto/#{File.basename(p, '.proto')}"
    FileUtils.mkdir_p dest_dir

    `protoc --go_out=#{proto_gen_go_path}/proto/#{File.basename(p, '.proto')} --proto_path=proto/  --proto_path=proto #{p}`
  end
end

