package snowflake

import (
	"crypto/sha256"
	"math/rand"
	"strconv"
	"time"
)

func SnowUID() (int64, error) {
	// Create a new Node with a Node number of 1
	if node, err := NewNode(1); err != nil {
		return 0, err
	} else {
		// Generate a snowflake ID.
		id := node.Generate()
		return id.Int64(), nil
	}
}

/*
RuleUID
@Description: package snowflake Generate the user id according to the custom rules.

步骤 1：获取时间戳
使用 time.Now().Format("20060102150405") 来获取当前精确到秒的时间戳字符串。
步骤 2：生成随机数和机器标识
通过 rand.Intn(9000) + 1000 生成一个范围在 1000 到 9999 的随机数，通过 rand.Intn(900) + 100 生成一个范围在 100 到 999 的机器标识。
步骤 3：组合数据并计算哈希
将时间戳、随机数和机器标识组合成一个字符串，然后对其进行 sha256 哈希计算。
最后，将得到的数字转换为字符串，并使用 zfill(12) 方法将其填充为 12 位长度，确保位数固定且全是数字。这样就得到了满足要求的全局唯一分布式用户 ID。这种方案结合了时间、随机数和哈希算法来增加随机性和唯一性。

@param machineLen: 机器标识范围,占3位长度
@return int64: 生成的用户id
*/
func RuleUID(machineLen int) int64 {
	timestamp := time.Now().Format("20060102150405")
	randomNum := rand.Intn(9000) + 1000
	machineID := rand.Intn(machineLen) + 100
	data := timestamp + strconv.Itoa(randomNum) + strconv.Itoa(machineID)
	hash := sha256.Sum256([]byte(data))
	userID := int64((uint64(hash[:6][0])<<32 | uint64(hash[:6][1])<<16 | uint64(hash[:6][2])<<8 | uint64(hash[:6][3])))
	return userID
}
