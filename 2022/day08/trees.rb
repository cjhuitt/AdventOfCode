class Node
  attr_reader :height, :row, :col
  attr_accessor :up, :down, :right, :left

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

  def visible?
    return true if edge?
    return true if upward.max < height
    return true if downward.max < height
    return true if rightward.max < height
    leftward.max < height
  end

  def scenic_score
    return 0 if edge?
    upward_vis * downward_vis * leftward_vis * rightward_vis
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

  def dirward_vis(&block)
    c = 0
    node = block.call(self)
    while node and node.height < height
      c += 1
      node = block.call(node)
    end
    c += 1 if node
    c
  end

  def upward_vis
    dirward_vis {|n| n.up}
  end

  def downward_vis
    dirward_vis {|n| n.down}
  end

  def rightward_vis
    dirward_vis {|n| n.right}
  end

  def leftward_vis
    dirward_vis {|n| n.left}
  end
end

nodes = []
last_row = []

File.foreach("input.txt", chomp: true).with_index do |line, row|
  heights = line.split("")
  row_nodes = []
  heights.each.with_index do |height, col|
    node = Node.new( row, col, height.to_i )
    node.connect_up( last_row[col] )
    node.connect_left( row_nodes.last )
    row_nodes << node
    nodes << node
  end
  last_row = row_nodes
end

visible = nodes.reject { |n| not n.visible? }
puts "Visible: #{visible.count}"

scenic_scores = nodes.collect { |n| n.scenic_score }
puts "Highest Scenic Score: #{scenic_scores.max}"
