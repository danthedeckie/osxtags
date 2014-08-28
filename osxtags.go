/*
 * OSX Tags
 * (C) 2014 Daniel Fairhead
 * MIT Licence
**/

package osxtags

import "fmt"

import "github.com/AlexSi/xattr"
import "howett.net/plist"

type TagColor int

const (
	None   TagColor = 0
	Gray   TagColor = 1
	Green  TagColor = 2
	Purple TagColor = 3
	Blue   TagColor = 4
	Yellow TagColor = 5
	Red    TagColor = 6
	Orange TagColor = 7
)

func (t TagColor) String() string {
	switch t {
	case None:
		return "None"
	case Gray:
		return "Gray"
	case Green:
		return "Green"
	case Purple:
		return "Purple"
	case Blue:
		return "Blue"
	case Yellow:
		return "Yellow"
	case Red:
		return "Red"
	case Orange:
		return "Orange"
	}
	return "unknown"
}

type ColorTag struct {
	Color TagColor
	Name  string
}

func GetColors(filename string) ([]ColorTag, error) {

	// pull out the raw metadata:

	colormeta, err := xattr.Get(filename, "com.apple.metadata:_kMDItemUserTags")

	if err != nil {
		return nil, err
	}

	// decode it from a binary plist:

	colList := []string{}

	_, err = plist.Unmarshal(colormeta, &colList)

	if err != nil {
		return nil, err
	}

	// split the stupid "\n" items into (color, name) pairs.

	toReturn := make([]ColorTag, len(colList), len(colList))

	for i, col := range colList {
		fmt.Sscanf(col, "%s\n%d", &toReturn[i].Name, &toReturn[i].Color)
	}

	return toReturn, nil

}
