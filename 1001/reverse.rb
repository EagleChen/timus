input = ""
$stdin.each_line do |line|
  input << line
end

input.scan(/\d+/).reverse.each do |str|
  puts "%.4f" % (str.to_f ** (1.0 / 2))
end
