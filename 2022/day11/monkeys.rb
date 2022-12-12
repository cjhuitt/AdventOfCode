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

  def inspect_items(monkey_list, relief_factor, cycle_factor, debug)
    @items.each do |item|
      val = inspect item
      val = val / relief_factor
      val = val % cycle_factor
      throw_to(monkey_list, val, item, debug)
    end
    @inspections += @items.length
    @items = []
  end

  def catch(item)
    @items << item
  end

  def apply_factor(factor)
    @items.each {|i| i.factor = factor}
  end

  private
  def extract_id(string)
    string.chop.split.last.to_i
  end

  def extract_items(string)
    string.split(":").last.split(", ").collect {|s| s.to_i}
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

  def throw_to(monkey_list, item, was, debug)
    rec = @recipients[item % @test == 0]
    monkey_list[rec].catch item
    if debug
      r = monkey_list[rec].id
      puts "\n#{@id} throwing #{item} (was #{was}) to #{r}"
      monkey_list.each {|m| puts "#{m}"}
    end
  end
end

monkeys1 = []
monkeys2 = []
input = ARGV.fetch(0, "input.txt")
lines = File.readlines(input, chomp: true)
until lines.empty?
  definition = lines.shift(7)
  monkeys1 << Monkey.new(definition.dup)
  monkeys2 << Monkey.new(definition.dup)
end

puts "#{monkeys1.length} total monkeys"
puts

def keep_away(monkeys, rounds, relief_factor)
  factor = relief_factor
  monkeys.each {|m| factor *= m.test}
  1.upto rounds do |round|
    monkeys.each do |monkey|
      monkey.inspect_items(monkeys, relief_factor, factor, false)
    end
  end
end

keep_away(monkeys1, 20, 3)
inspections = monkeys1.collect {|m| m.inspections }.sort
puts "Part 1 monkey business: #{inspections.pop * inspections.pop}"

$stdout.sync = true
100.downto 1 do |n|
  keep_away(monkeys2, 100, 1)
  print "."
  print "\b\b\b\b\b#    \b\b\b\b" if n % 5 == 0
end
inspections = monkeys2.collect {|m| m.inspections }.sort
puts "\nPart 2 monkey business: #{inspections.pop * inspections.pop}"

