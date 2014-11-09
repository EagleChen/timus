num = gets.to_i

db = []
num.times { db << gets.to_i }
gets

db.sort!
num = gets.to_i
num.times do
  index = gets.to_i
  puts db[index-1]
end
