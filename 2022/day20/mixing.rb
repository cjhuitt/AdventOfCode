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
    n = @head
    n = n.next until n == node
    v = n.value # % @length - doesn't work how I want on negatives
    if v < 0
      (v...0).each { move_prev(n) }
      #backward
    else
      #forward
      (0...v).each { move_next(n) }
    end
  end

  def next(node, steps)
    (0...steps).each do
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
  def move_prev(node)
    if node == @head
      # N@H, m, ...y, z@T => m@h, ..., y, N, z@T
      m = node.next
      z = @tail
      y = z.prev
      m.prev = nil
      y.next = node
      node.prev = y
      node.next = z
      z.prev = node
      @head = m
    else
      # l, m, N, o => l, N, m, o
      m = node.prev
      l = m.prev
      o = node.next
      l.next = node unless l.nil?
      node.prev = l
      node.next = m
      m.prev = node
      m.next = o
      o.prev = m unless o.nil?
      @head = node if @head == m
      @tail = m if @tail == node
    end
  end

  def move_next(node)
    if node == @tail
      # a@H, b, ... m, N@T => a@H, N, b, ... m@T
      a = @head
      b = a.next
      m = node.prev
      m.next = nil
      a.next = node
      node.prev = a
      node.next = b
      b.prev = node
      @tail = m
    else
      # m, N, o, p => m, o, N, p
      m = node.prev
      o = node.next
      p = o.next
      m.next = o unless m.nil?
      o.prev = m
      o.next = node
      node.prev = o
      node.next = p
      p.prev = node unless p.nil?
      @tail = node if @tail == o
      @head = o if @head == node
    end
  end
end

order = []
list = List.new
input = ARGV.fetch(0, "input.txt")
File.foreach(input, chomp: true).with_index do |line, index|
  v = line.to_i
  order << v
  list.add(Struct::Node.new(v, nil, nil))
end

order.each do |val|
  list.move(list.find(val))
#  puts "#{val}: #{list}"
end

index = list.find(0)
values = (0..2).collect do
  index = list.next(index, 1000)
  index.value
end
puts "#{values} sum to #{values.sum}"
