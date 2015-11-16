require 'nokogiri'
require 'httparty'


IMAGE_DIR = "./images"
GDS_DIR = File.expand_path(File.join(File.dirname(__FILE__), "..", "gds"))

TYPE_INDEX_MAP = %w(building defensive_building unit resource)

class Crawler
  include HTTParty
  base_uri "http://www.boombeachhq.com"

  def self.index
    @page ||= Nokogiri::HTML get("/wiki/").body
    @page.css(".bb-unit-infocolumn")
  end

  def self.base_building
    do_the_craw(0)
  end

  def self.defensive_building
    do_the_craw(1)
  end

  def self.unit
    do_the_craw(2)
  end

  def self.resource
    do_the_craw(3)
  end

  def self.do_the_craw(i)
    links = index[i].css("a")
    links.each do |building|
      save_gds TYPE_INDEX_MAP[i], building.attr("href")
      image_src = building.css("img").attr("src")
      save_image(image_src)
    end
  end

  def self.save_image src
    `curl #{src} > ./images/#{File.basename(src).sub(/\?(.*)$/, "")}`
  end

  def self.save_gds type, src
    file_name = File.basename(File.basename(src).sub(/\?(.*)$/, ""), ".png")
    p = Nokogiri::HTML get(src).body
    csv_content = ""
    table = p.css(".table-responsive")
    table.css("tr").each_with_index do |tr, index|
      if index.zero?
        csv_content << tr.css("td").map {|td| normalize_header(td.text) }.join(",")
        csv_content << "\n"
      else
        csv_content << tr.css("td").map {|td| td.text }.join(",")
        csv_content << "\n"
      end
    end
    puts csv_content
    File.open(File.join(GDS_DIR, type, file_name) + ".csv", "w") do |f|
     f << csv_content
    end
  end

  def self.normalize_header(header)
    header.gsub(" ", "_").downcase
  end
end

Crawler.base_building
Crawler.unit
Crawler.resource
Crawler.defensive_building
