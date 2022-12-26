Struct.new('Node', :value, :prev, :next)

class List
  attr_accessor :length
  def initialize
    @head = @tail = nil
    @length = 0
  end

  def add(node)
    @head = node if @head.nil?
    if @tail.nil?
      @tail = @head
    else
      @tail.next = node
      node.prev = @tail
      @tail = node
    end
    @length += 1
  end

  def find(v)
    node = @head
    until node.nil?
      return node if node.value == v
      node = node.next
    end
  end

  def move(node)
    return if node.value == 0
    after = advance(node, node.value)
    return if after == node
    remove(node)
    insert(node, after)
  end

  def remove(node)
    p = node.prev
    n = node.next
    p.next = n unless p.nil?
    n.prev = p unless n.nil?
    @head = n if @head == node
    @tail = p if @tail == node
  end

  def insert(node, after)
    if after.nil?
      @tail.next = node
      node.prev = @tail
      @tail = node
    else
      p = after
      n = after.next
      p.next = node
      node.prev = p
      node.next = n
      n.prev = node unless n.nil?
      @tail = node if @tail == p
    end
  end

  def advance(node, steps)
    normed = normalize(steps)
    (0...normed).each do
      if node == @tail
        node = @head
      else
        node = node.next
      end
    end
    node
  end

  def to_s
    n = @head
    lines = [n.value]
    until n == @tail
      n = n.next
      lines << n.value
    end
    lines.join(", ")
  end

  private
  def normalize(val)
    r = val.abs % @length
    return r if val >= 0
    @length - r - 1
  end
end

order = []
list = List.new
zero = nil
input = ARGV.fetch(0, "input.txt")
File.foreach(input, chomp: true).with_index do |line, index|
  v = line.to_i
  order << v
  node = Struct::Node.new(v, nil, nil)
  list.add(node)
  zero = node if v == 0
end

order.each.with_index do |val, index|
  list.move(list.find(val))
#  puts "#{val}: #{list}"
end

index = zero
values = (0..2).collect do
  index = list.advance(index, 1000)
  index.value
end
puts "#{values} sum to #{values.sum}"
