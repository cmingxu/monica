task :gen do
  proto_file_path = "protobuf_defs"
  proto_gen_go_path = "protogos"

  Dir.chdir "protobuf_defs"
  Dir["*.proto"].each do |p|
    dest_dir = File.expand_path("../#{proto_gen_go_path}/#{File.basename(p, '.proto')}")
    FileUtils.mkdir_p dest_dir
    puts "protoc --go_out=../#{proto_gen_go_path}/#{File.basename(p, '.proto')}/ #{p}"
    `protoc --go_out=../#{proto_gen_go_path}/#{File.basename(p, '.proto')}/ #{p}`
  end
  Dir.chdir ".."
end

