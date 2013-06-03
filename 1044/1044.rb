#how i create the "result" hash
#def cal_sum(num)
#  sum = 0
#  current = -1
#  loop do
#    sum += num % 10
#    num = num / 10
#    break if num == 0
#  end
#  sum
#end
#
#def create_record(bits)
#  record = {}
#  record.default = 0
#  max = 9 * bits
#  max_num = 10 ** bits - 1
#  (0..max_num).each { |i| record[cal_sum(i)] += 1 }
#  record
#end
#
#def analyze_num(bits)
#  record = create_record(bits / 2)
#  record.values.map { |i| i*i }.reduce(:+)
#end
#
#result = {}
#2.step(8, 2) { |i| result[i] = analyze_num(i) }
#puts result
#end :how i create the "result" hash

result = {2=>10, 4=>670, 6=>55252, 8=>4816030}

puts result[gets.to_i]
