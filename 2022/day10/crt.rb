class CPU
  attr_reader :regx, :clock, :signal_strength

  def initialize()
    @regx = 1
    @clock = 0
    @signal_strength = []
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
    accumulate_signal() if [20, 60, 100, 140, 180, 220].include? @clock
#    puts "Tick #{@clock}: #{@regx}"
  end

  def accumulate_signal
    @signal_strength << @regx * @clock
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

puts "Summed strengths: #{cpu.signal_strength.sum}"
