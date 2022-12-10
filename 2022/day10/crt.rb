class CPU
  attr_reader :regx, :regx_history, :clock

  def initialize()
    @regx = 1
    @regx_history = []
    @clock = 0
  end

  def noop()
    tick
  end

  def addx(val)
    tick
    tick
    @regx += val
  end

  private
  def tick
    @clock += 1
    @regx_history << @regx
  end
end

cpu = CPU.new
File.foreach("input.txt", chomp: true) do |line|
  instruction = line.split
  case instruction[0]
  when "noop"
    cpu.noop
  when "addx"
    cpu.addx instruction[1].to_i
  else
    raise "Invalid instruction"
  end
end

signal_strength = [20, 60, 100, 140, 180, 220].collect do |tick|
  cpu.regx_history[tick - 1] * tick
end
puts "Summed strengths: #{signal_strength.sum}"


