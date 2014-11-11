num = gets.to_i

result = Array.new(101) { [] }
num.times do |index|
  row = gets
  m = row.split(" ")[1].to_i
  result[m] << row
end

100.downto(0).each do |i|
  result[i].each do |team|
    print team
  end
end
