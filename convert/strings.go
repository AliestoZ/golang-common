package convert

func PStringToString(pString *string) string {
	if pString != nil {
		return *pString
	}
	return ""
}
