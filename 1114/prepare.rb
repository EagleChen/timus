arr = Array.new(16) { |index| Array.new(16, 0) }

arr[0][1] = 1
(1..15).each do |i|
  arr[i][0] = 1
  (1..i).each do |j|
    (1..j).each do |k|
      puts "#{i}, #{j}, #{k}"
      arr[i][j] += arr[i-j][k]
    end
    puts "arr[#{i}][#{j}] = #{arr[i][j]}"
  end
end

puts "["
arr.each do |a|
  puts "[#{a.join(", ")}],"
end
puts "]"
