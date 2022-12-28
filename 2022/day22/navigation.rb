class Node
  attr_reader :row, :col, :type
  attr_accessor :up, :down, :right, :left

  def initialize( r, c, type )
    @row = r
    @col = c
    @type = type
  end

  def to_s
    "Node(#{row},#{col}): #{@type}"
  end

  def connect_up( node )
    @up = node
    node.down = self if node
  end

  def connect_left( node )
    @left = node
    node.right = self if node
  end

  def move(dir)
    n = [@right, @down, @left, @up][dir]
    if n.nil?
      wrap = [@left, @up, @right, @down][dir]
      until wrap.nil?
        n = wrap
        wrap = [n.left, n.up, n.right, n.down][dir]
      end
    end
    return self if n.type == '#'
    n
  end
end

nodes = []
last_row = []
directions = nil
input = ARGV.fetch(0, "input.txt")
File.foreach(input, chomp: true).with_index do |line, row|
  if line.match? /^\d/
    directions = line
    next
  end
  row_nodes = []
  line.chars.each_with_index do |ch, col|
    node = nil
    if ch != ' '
      node = Node.new(row, col, ch)
      node.connect_up( last_row[col] )
      node.connect_left( row_nodes.last )
    end
    row_nodes << node
    nodes << node
  end
  last_row = row_nodes
end
nodes.compact!

def change_dir(ch, current)
  (current + (ch == "R" ? 1 : 3)) % 4
end

pos = nodes.first
dir = 0
directions.scan(/(?<steps>\d+)(?<turn>[RL])?/) do |steps, turn|
  1.upto(steps.to_i).each {pos = pos.move(dir)}
  dir = change_dir(turn, dir) unless turn.nil?
end

puts "Password is #{(pos.row + 1) * 1000 + (pos.col + 1) * 4 + dir}"
