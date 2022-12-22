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
    faces_to nil
  end

  def faces_to(val)
    [@up, @down, @right, @left, @fore, @back].keep_if {|f| f == val}.length
  end
end

total = 0
input = ARGV.fetch(0, "input.txt")
lava_nodes = {}
File.foreach(input, chomp: true).with_index do |line, index|
  coords = line.split(",")
  node = Node.new(coords[0].to_i, coords[1].to_i, coords[2].to_i)
  lava_nodes[node.point] = node
end

lava_nodes.each_value do |node|
  node.connect_left lava_nodes[Struct::Point.new(node.point.x - 1, node.point.y, node.point.z)]
  node.connect_up lava_nodes[Struct::Point.new(node.point.x, node.point.y + 1, node.point.z)]
  node.connect_fore lava_nodes[Struct::Point.new(node.point.x, node.point.y, node.point.z + 1)]
end

open = lava_nodes.each_value.collect {|n| n.open_faces}.sum
puts "#{lava_nodes.size} nodes, #{open} open faces"

first = lava_nodes.values.first
min_x = max_x = first.point.x
min_y = max_y = first.point.y
min_z = max_z = first.point.z

lava_nodes.each_value do |node|
  min_x = [min_x, node.point.x].min
  max_x = [max_x, node.point.x].max
  min_y = [min_y, node.point.y].min
  max_y = [max_y, node.point.y].max
  min_z = [min_z, node.point.z].min
  max_z = [max_z, node.point.z].max
end

min_x -= 1
max_x += 1
min_y -= 1
max_y += 1
min_z -= 1
max_z += 1

puts "X: #{min_x}-#{max_x}"
puts "Y: #{min_y}-#{max_y}"
puts "Z: #{min_z}-#{max_z}"

def around(point)
  x = point.x
  y = point.y
  z = point.z
  [Struct::Point.new(x - 1, y, z), Struct::Point.new(x + 1, y, z),
   Struct::Point.new(x, y - 1, z), Struct::Point.new(x, y + 1, z),
   Struct::Point.new(x, y, z - 1), Struct::Point.new(x, y, z + 1)]
end

def in_bounds(point)
  (min_x..max_x).include? point.x and
    (min_y..max_y).include? point.y and
    (min_z..max_z).include? point.z
end

directions = [:@left, :@right, :@down, :@up, :@back, :@fore]

infinity = true
processed = []
queue = [Struct::Point.new(min_x, min_y, min_z)]
until queue.empty?
  current = queue.pop
  processed << current
  around(current).each_with_index do |point, index|
    next if not (min_x..max_x).include? point.x
    next if not (min_y..max_y).include? point.y
    next if not (min_z..max_z).include? point.z
    next if processed.include? point
    if lava_nodes.has_key? point
      lava_nodes[point].instance_variable_set(directions[index], infinity)
    else
      queue.push point
    end
  end
end

external = lava_nodes.each_value.collect {|n| n.faces_to infinity}.sum
puts "#{external} External faces"
