class Recipe
  attr_reader :ore, :clay, :obsidian

  def initialize(ore, clay, obs)
    @ore = ore
    @clay = clay
    @obsidian = obs
  end

  def can_run?(current_resources)
    current_resources[:ore] >= @ore and
      current_resources[:clay] >= @clay and
      current_resources[:obsidian] >= @obsidian
  end

  def consume(resources)
    resources[:ore] -= @ore
    resources[:clay] -= @clay
    resources[:obsidian] -= @obsidian
    resources
  end

  def to_s
    parts = []
    parts << "#{@ore} ore" if @ore > 0
    parts << "#{@clay} clay" if @clay > 0
    parts << "#{@obsidian} obsidian" if @obsidian > 0
    parts.join(", ")
  end
end

class RecipeSet
  attr_reader :id

  def initialize(line)
    @line = line
    parts = line.split(":")
    @id = parse_id(parts[0])
    costs = parts[1].split(". ")
    @ore = parse_recipe(costs[0])
    @clay = parse_recipe(costs[1])
    @obsidian = parse_recipe(costs[2])
    @geode = parse_recipe(costs[3])
  end

  def to_s
    "Blueprint #{@id}: New ore is #{@ore}; new clay is #{@clay}; new obsidian is #{@obsidian}; new geode is #{@geode}"
  end

  def max_geodes(steps, resources, robots)
    return resources if steps <= 0
    options = []
    if @ore.can_run? resources
      new_resources = @ore.consume resources.dup
      robots.each {|r| new_resources[r] += 1}
      new_robots = robots.dup
      new_robots << :ore
      options << max_geodes(steps - 1, new_resources, new_robots)
    end
    if @clay.can_run? resources
      new_resources = @clay.consume resources.dup
      robots.each {|r| new_resources[r] += 1}
      new_robots = robots.dup
      new_robots << :clay
      options << max_geodes(steps - 1, new_resources, new_robots)
    end
    if @obsidian.can_run? resources
      new_resources = @obsidian.consume resources.dup
      robots.each {|r| new_resources[r] += 1}
      new_robots = robots.dup
      new_robots << :obsidian
      options << max_geodes(steps - 1, new_resources, new_robots)
    end
    if @geode.can_run? resources
      new_resources = @geode.consume resources.dup
      robots.each {|r| new_resources[r] += 1}
      new_robots = robots.dup
      new_robots << :geode
      options << max_geodes(steps - 1, new_resources, new_robots)
    end
    if true # always test just getting more resources
      new_resources = resources.dup
      robots.each {|r| new_resources[r] += 1}
      options << max_geodes(steps - 1, new_resources, robots)
    end
    max = options.first
    options.each do |o|
      max = o if max[:geode] < o[:geode]
    end
    max
  end

  private
  def parse_id(str)
    str.match(/Blueprint (?<id>\d+)/) {|m| m[:id].to_i }
  end
  def parse_recipe(str)
    regex = /((?<ore>\d+) ore)( and (?<clay>\d+) clay)?( and (?<obsidian>\d+) obsidian)?/
    str.match(regex) {|m| Recipe.new(m[:ore].to_i, m[:clay].to_i, m[:obsidian].to_i) }
  end
end

recipes = []
input = ARGV.fetch(0, "input.txt")
File.foreach(input, chomp: true).with_index do |line, index|
  recipes << RecipeSet.new(line)
end

puts "#{recipes.length} lines"
puts "#{recipes.first.max_geodes(24, Hash.new(0), [:ore])}"
#scores = recipes.collect {|r| r.id * r.max_geodes(24)}
#puts "Scores: #{scores} sum to #{scores.sum}"
