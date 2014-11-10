arr1, arr2 = [], []
num = gets.to_i
num.times { arr1 << gets.to_i }
num = gets.to_i
num.times { arr2 << gets.to_i }

i, j = 0, 0
loop do
  sum = arr1[i] + arr2[j]
  if sum > 10000
    j += 1
    if j >= arr2.size
      puts "NO"
      exit
    end
  elsif sum == 10000
    puts "YES"
    exit
  else
    i += 1
    if i >= arr1.size
      puts "NO"
      exit
    end
  end
end
