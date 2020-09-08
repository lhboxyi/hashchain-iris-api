package util_common

import (
	"encoding/base64"
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/core"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/cryptosuite"
	"iris-api/com/util"
	"math"
)

type MerkleUtil struct {
}

func (mu *MerkleUtil) GetMerkleRoot(hashValueList util.ArrayList, suite core.CryptoSuite)  {
	treeDepth := math.Log(float64(hashValueList.Size())) / math.Log(2)
	for j := 0; j < int(treeDepth); j++ {
		var halfhashValueList util.ArrayList
		halfLength := int(math.Ceil(float64(hashValueList.Size()) / 2.0))
		for i := 0; i < halfLength; i++ {
			if i*2+1 == hashValueList.Size() {
				halfhashValueList.Add(hashValueList.Get(2 * i))
			} else {
				///hashValue :=
			}
		}
	}
}

func (mu *MerkleUtil) computeHashValueByLeaf(leftHashValue string, rightHashValue string, suite core.CryptoSuite) string {
	totalhashValue := leftHashValue + rightHashValue
	hashBytes, err := suite.Hash([]byte(totalhashValue),cryptosuite.GetSHA256Opts())
	if err != nil {
		fmt.Println(err)
	}
	var hashValue string = base64.NewEncoding("").EncodeToString(hashBytes)
	return hashValue
}

func computeMerkleTree(hashValueList util.ArrayList,suite core.CryptoSuite)  {
	var merkleTree util.ArrayListContainArrayList
	treeDepth := int(math.Log(float64(hashValueList.Size())/2.0))
	for j:=0;j<treeDepth;j++{
		merkleTree.Add(hashValueList)
		var halfHashValueList util.ArrayList
		halfLength := int(math.Ceil(float64(hashValueList.Size())/2.0))
		for i:=0;i<halfLength;i++ {
			if i*2+1 == hashValueList.Size(){
				halfHashValueList.Add(hashValueList.Get(2*i))
			}else{
				//hashValue :=
			}
		}
	}
}
