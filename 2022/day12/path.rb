height = {
  "a" =>  1, "b" =>  2, "c" =>  3, "d" =>  4, "e" =>  5, "f" =>  6, "g" =>  7,
  "h" =>  8, "i" =>  9, "j" => 10, "k" => 11, "l" => 12, "m" => 13, "n" => 14,
  "o" => 15, "p" => 16, "q" => 17, "r" => 18, "s" => 19, "t" => 20, "u" => 21,
  "v" => 22, "w" => 23, "x" => 24, "y" => 25, "z" => 26, "S" =>  1, "E" => 26
}

class Node
  attr_reader :height, :row, :col
  attr_accessor :up, :down, :right, :left
  attr_accessor :allowed_up, :allowed_down, :allowed_right, :allowed_left
  attr_accessor :distance

  def initialize( r, c, h )
    @row = r
    @col = c
    @height = h
  end

  def to_s
    "Node(#{row},#{col}): #{height}"
  end

  def connect_up( node )
    @up = node
    node.down = self if node
  end

  def connect_left( node )
    @left = node
    node.right = self if node
  end

  def prune_connections
    @allowed_up = @up if @up and @up.height <= @height + 1
    @allowed_down = @down if @down and @down.height <= @height + 1
    @allowed_left = @left if @left and @left.height <= @height + 1
    @allowed_right = @right if @right and @right.height <= @height + 1
  end

  def begin_search
    update_distance(0)
  end

  def reduce_distance_to(node)
    if @up == node and @allowed_up and (not distance or node.distance + 1 < distance)
      update_distance(node.distance + 1)
    elsif @down == node and @allowed_down and (not distance or node.distance + 1 < distance)
      update_distance(node.distance + 1)
    elsif @left == node and @allowed_left and (not distance or node.distance + 1 < distance)
      update_distance(node.distance + 1)
    elsif @right == node and @allowed_right and (not distance or node.distance + 1 < distance)
      update_distance(node.distance + 1)
    end
  end

  private
  def update_distance(distance)
    @distance = distance
    [@up, @down, @left, @right].compact.each {|n| n.reduce_distance_to(self)}
  end
end

total = 0

nodes = []
last_row = []
start_node = nil
end_node = nil
input = ARGV.fetch(0, "input.txt")
File.foreach(input, chomp: true).with_index do |line, row|
  heights = line.split("")
  row_nodes = []
  heights.each.with_index do |h, col|
    node = Node.new( row, col, height[h] )
    node.connect_up( last_row[col] )
    node.connect_left( row_nodes.last )
    row_nodes << node
    nodes << node
    start_node = node if h == "S"
    end_node = node if h == "E"
  end
  last_row = row_nodes
end

nodes.each {|n| n.prune_connections}
end_node.begin_search
puts "\nDistance from start is #{start_node.distance}"
