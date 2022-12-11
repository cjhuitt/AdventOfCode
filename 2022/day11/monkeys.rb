class Num
  attr_reader :val
  def initialize(val)
    @val = val
  end

  def +(other)
    if other.is_a? Num
      Num.new(@val + other.val)
    else
      Num.new(@val + other)
    end
  end

  def *(other)
    if other.is_a? Num
      Num.new(@val * other.val)
    else
      Num.new(@val * other)
    end
  end

  def /(other)
    if other.is_a? Num
      Num.new((@val / other.val).floor)
    else
      Num.new((@val / other).floor)
    end
  end

  def divisible(other)
    @val % other == 0
  end

  def to_s
    "#{@val}"
  end
end

class Monkey
  attr_reader :id, :test, :inspections

  def initialize(lines)
    @id = extract_id lines.shift
    @items = extract_items lines.shift
    @operation = extract_op lines.shift
    @test = extract_test lines.shift
    @recipients = {true => extract_recipient(lines.shift),
                   false => extract_recipient(lines.shift)}
    @inspections = 0
  end

  def to_s
    "Monkey #{@id} (#{@inspections} inspections) has items: #{@items.collect{|i| i.to_s}.join(", ")}"
  end

  def inspect_items(monkey_list)
    @items.each do |item|
      val = inspect item
      val = val / 3
      throw_to(monkey_list, val)
    end
    @inspections += @items.length
    @items = []
  end

  def catch(item)
    @items << item
  end

  private
  def extract_id(string)
    string.chop.split.last.to_i
  end

  def extract_items(string)
    string.split(":").last.split(", ").collect {|s| Num.new(s.to_i)}
  end

  def extract_op(string)
    string.split("= old ").last.split
  end

  def extract_test(string)
    string.split.last.to_i
  end

  def extract_recipient(string)
    string.split.last.to_i
  end

  def inspect(item)
    case @operation.first
    when "*"
      return item * item if @operation.last == "old"
      item * @operation.last.to_i
    when "+"
      return item + item if @operation.last == "old"
      item + @operation.last.to_i
    else
      item
    end
  end

  def throw_to(monkey_list, item)
    monkey_list[@recipients[item.divisible(test)]].catch item
  end
end

monkeys = []
input = ARGV.fetch(0, "input.txt")
lines = File.readlines(input, chomp: true)
until lines.empty?
  definition = lines.shift(7)
  monkeys << Monkey.new(definition.dup)
end

puts "#{monkeys.length} total monkeys"
puts

def keep_away(monkeys, rounds)
  1.upto rounds do |round|
    monkeys.each do |monkey|
      monkey.inspect_items(monkeys)
    end
    #puts "\nRound #{round}"
    #monkeys.each {|m| puts "#{m}"}
  end
end

keep_away(monkeys, 20)
inspections = monkeys.collect {|m| m.inspections }.sort
puts "Part 1 monkey business: #{inspections.pop * inspections.pop}"

