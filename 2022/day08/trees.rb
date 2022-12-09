class Node
  attr_reader :height
  attr_accessor :up, :down, :right, :left

  def initialize( h )
    @height = h
    @up = nil
    @down = nil
    @right = nil
    @left = nil
  end

  def connect_up( node )
    @up = node
    node.down = self if node
  end

  def connect_left( node )
    @left = node
    node.right = self if node
  end

  def visible?
    return true if edge?
    return true if upward.max < height
    return true if downward.max < height
    return true if rightward.max < height
    leftward.max < height
  end

  private
  def edge?
    @up == nil or @down == nil or @right == nil or @left == nil
  end

  def dirward(&block)
    set = []
    node = block.call(self)
    while node
      set << node.height
      node = block.call(node)
    end
    set
  end

  def upward
    dirward {|n| n.up}
  end

  def downward
    dirward {|n| n.down}
  end

  def rightward
    dirward {|n| n.right}
  end

  def leftward
    dirward {|n| n.left}
  end
end

nodes = []
last_row = []

File.foreach("input.txt", chomp: true) do |line|
  heights = line.split("")
  row = []
  heights.each.with_index do |height, index|
    node = Node.new( height.to_i )
    node.connect_up( last_row[index] )
    node.connect_left( row.last )
    row << node
    nodes << node
  end
  last_row = row
end

visible = nodes.reject { |n| not n.visible? }
puts "Visible: #{visible.count}"

