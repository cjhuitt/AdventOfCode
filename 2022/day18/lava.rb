Struct.new('Point', :x, :y, :z)

class Node
  attr_reader :point
  attr_accessor :up, :down, :right, :left, :fore, :back

  def initialize( x, y, z )
    @point = Struct::Point.new(x, y, z)
  end

  def to_s
    "Node(#{x},#{y}, #{z})"
  end

  def connect_up( node )
    @up = node
    node.down = self if node
  end

  def connect_left( node )
    @left = node
    node.right = self if node
  end

  def connect_fore( node )
    @fore = node
    node.back = self if node
  end

  def open_faces
    6 - [@up, @down, @right, @left, @fore, @back].compact.length
  end
end

total = 0
input = ARGV.fetch(0, "input.txt")
nodes = {}
File.foreach(input, chomp: true).with_index do |line, index|
  coords = line.split(",")
  node = Node.new(coords[0].to_i, coords[1].to_i, coords[2].to_i)
  nodes[node.point] = node
end

nodes.each_value do |node|
  node.connect_left nodes[Struct::Point.new(node.point.x - 1, node.point.y, node.point.z)]
  node.connect_up nodes[Struct::Point.new(node.point.x, node.point.y + 1, node.point.z)]
  node.connect_fore nodes[Struct::Point.new(node.point.x, node.point.y, node.point.z + 1)]
end

open = nodes.each_value.collect {|n| n.open_faces}.sum
puts "#{nodes.size} nodes, #{open} open faces"
