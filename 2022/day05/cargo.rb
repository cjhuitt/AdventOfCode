stacks1 = [Array.new, Array.new, Array.new, Array.new, Array.new,
           Array.new, Array.new, Array.new, Array.new]
stacks2 = [Array.new, Array.new, Array.new, Array.new, Array.new,
           Array.new, Array.new, Array.new, Array.new]
initializing = true

File.foreach("input.txt").with_index do |line, index|
  unless line.strip.empty?
    if initializing
      vals = line.scan /.(.)...(.)...(.)...(.)...(.)...(.)...(.)...(.)...(.)./m
      puts "#{vals}"
      vals.flatten.each.with_index do |val, index|
        stacks1[index].push val if not val.strip.empty?
        stacks2[index].push val if not val.strip.empty?
      end
    else
      a, b, c = line.scan /(\d+)/
      count, from, to = a.join.to_i, b.join.to_i - 1, c.join.to_i - 1
      shift1 = stacks1[from].pop count
      shift2 = stacks2[from].pop count
      shift1.reverse.each do |crate|
        stacks1[to].push crate
      end
      shift2.each do |crate|
        stacks2[to].push crate
      end
    end
  else
    initializing = false
    stacks1.each do |stack|
      stack.reverse!
    end
    stacks2.each do |stack|
      stack.reverse!
    end
  end
end

puts "First Tops: #{stacks1.map{|s| s.last}.join}"
puts "Second Tops: #{stacks2.map{|s| s.last}.join}"
