n = gets.to_i

nums = {}
nums.default = 0
max = 0
max_num = "-1"
half = n / 2
n.times do
  num = gets
  nums[num] += 1
  if (nums[num] > max)
    max = nums[num]
    max_num = num
  end
  break if max > half
end

printf max_num
