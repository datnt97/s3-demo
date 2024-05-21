package unicode

import "strings"

var VnToLatin = map[rune]string{
	'à': "a", 'á': "a", 'ả': "a", 'ã': "a", 'ạ': "a",
	'ầ': "a", 'ấ': "a", 'ẩ': "a", 'ẫ': "a", 'ậ': "a",
	'è': "e", 'é': "e", 'ẻ': "e", 'ẽ': "e", 'ẹ': "e",
	'ề': "e", 'ế': "e", 'ể': "e", 'ễ': "e", 'ệ': "e",
	'ì': "i", 'í': "i", 'ỉ': "i", 'ĩ': "i", 'ị': "i",
	'ò': "o", 'ó': "o", 'ỏ': "o", 'õ': "o", 'ọ': "o",
	'ồ': "o", 'ố': "o", 'ổ': "o", 'ỗ': "o", 'ộ': "o",
	'ờ': "o", 'ớ': "o", 'ở': "o", 'ỡ': "o", 'ợ': "o",
	'ù': "u", 'ú': "u", 'ủ': "u", 'ũ': "u", 'ụ': "u",
	'ừ': "u", 'ứ': "u", 'ử': "u", 'ữ': "u", 'ự': "u",
	'ỳ': "y", 'ý': "y", 'ỷ': "y", 'ỹ': "y", 'ỵ': "y",
	'đ': "d", 'ê': "e", 'ă': "a", 'ơ': "o", 'ắ': "a",
	'â': "a", 'ư': "u", 'ô': "o", 'Ô': "o", 'Ồ': "o",
	'Ố': "o", 'Ổ': "o", 'Ỗ': "o", 'Ộ': "o",
}

func ToLatin(input string) string {
	var result strings.Builder
	for _, char := range input {
		if replacement, exists := VnToLatin[char]; exists {
			result.WriteString(replacement)
		} else {
			result.WriteRune(char)
		}
	}
	return result.String()
}
