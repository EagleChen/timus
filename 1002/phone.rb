def translate(origin)
  origin.split(//).map do |c|
    case c
    when /[ij]/ then "1"
    when /[abc]/ then "2"
    when /[def]/ then "3"
    when /[gh]/ then "4"
    when /[kl]/ then "5"
    when /[mn]/ then "6"
    when /[prs]/ then "7"
    when /[tuv]/ then "8"
    when /[wxy]/ then "9"
    when /[oqz]/ then "0"
    end
  end.join
end

Node = Struct.new(:origin, :value, :children) do
  def add_child(originWord, word, index)
    key = word[index]
    unless children.key?(key)
      children[key] = Node.new(nil, nil, {})
    end

    if (index == word.length - 1)
      children[key].origin = originWord
      children[key].value = word
    else
      children[key].add_child(originWord, word, index+1)
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
#now = Time.now
  num.times do
    origin = gets.chomp
    word = translate(origin)
    root.add_child(origin, word, 0)
  end
#p "tree completed : #{Time.now - now}"

  puts search(root, phone)
end
