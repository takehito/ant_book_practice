var (
	used []bool
	perm []int

	// 重複があっても全ての順列を生成する
	// next_permutationは辞書順で次の順列を生成する
	perm2 []int
)

// {0,1,2,3,4, ... , n-1}の並び替えn!通りを生成する

func permutation(pos int, n int)  {
	if pos == n {
		// permに対する操作
		return
	}

	// permのpos番目を0~n-1のどれにするかのループ
	for i := 0; i < n; i++ {
		if !used[i] {
			perm[pos] = i
			// iを使ったのでフラグをtrueにしておく
			used[i] = true
			permutation1(pos + 1, n)
			// 戻ってきたらフラグを戻しておく
			used[i] = false
		}
	}
	return	
}

func permutation2(n int) {
	for i := 0; i < n; i++ {
		perm2[i] = i;
	}
	for {/* 
		perm2に対する操作
	*/
	if next_permutation(perm2, perm2 + n)) {
		break
	}
	}
	// 全ての順列を生成し終わったらnext_permutationはfalseを返す
	return
}