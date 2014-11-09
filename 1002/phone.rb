def translate(origin)
  dict = { "a" => "2", "b" => "2", "c" => "2", "d" => "3", "e" => "3", "f" => "3",
           "g" => "4", "h" => "4", "i" => "1", "j" => "1", "k" => "5", "l" => "5",
           "m" => "6", "n" => "6", "o" => "0", "p" => "7", "q" => "0", "r" => "7",
           "s" => "7", "t" => "8", "u" => "8", "v" => "8", "w" => "9", "x" => "9",
           "y" => "9", "z" => "0" }
  origin.chars.map { |c| dict[c] }.join
end

class Node
  attr_accessor :origin, :value, :children

  def initialize(o, v, c)
    @origin = o
    @value = v
    @children = c
  end

  def add_child(originWord, word, index)
    candidates = @children
    loop do
      key = word[index]
      unless candidates.key?(key)
        candidates[key] = Node.new(nil, nil, {})
      end

      if (index == word.length - 1)
        candidates[key].origin = originWord
        candidates[key].value = word
        break
      else
        candidates = candidates[key].children
        index += 1
      end
    end
  end

  # debug info
  def to_s
    res = "#{value} #{len}\n"
    children.each do |k, v|
      res << "#{k} -> #{v.to_s}"
    end
    res
  end
end

def search_with_prefix(root, phone, index)
  candidates = root.children
  words = []
  while index < phone.length
    key = phone[index]
    if candidates.key?(key)
      words << candidates[key].origin if candidates[key].value
      index += 1
      candidates = candidates[key].children
    else
      break
    end
  end
  words.reverse
end

def search(root, phone)
  prefixArr = [[]]
  visited = {}
  while !prefixArr.empty?
    prefixes = prefixArr.shift
    idx = prefixes.map(&:length).reduce(&:+) || 0
    # don't waste time on visited
    if visited[idx]
      next
    else
      visited[idx] = 1
    end
    return prefixes.join(" ") if idx == phone.length
    wordsToAdd = search_with_prefix(root, phone, idx)
    unless wordsToAdd.empty?
      prefixArr.concat(wordsToAdd.map { |word| prefixes + [word] })
    end
  end
  "No solution."
end

loop do
  phone = gets.chomp
  break if phone == "-1"

  root = Node.new(nil, nil, {})
  num = gets.to_i
now = Time.now
  num.times do
    origin = gets.chomp
    word = translate(origin)
    root.add_child(origin, word, 0)
  end
p "tree completed : #{Time.now - now}"

  puts search(root, phone)
end
