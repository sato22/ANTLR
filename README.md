# ANTLR

### プログラム
```
	#include <stdio.h>
  
	void main()
	{
 
	  int a = 3;
	  int b = 4;
	 
	  printf("a+b = %d\n",a+b);
	}
```

### 実行結果
```
b175949@trumpet:~/parsar$ go test
2020/11/24 03:48:00 EnterEveryRule
2020/11/24 03:48:00 EnterDeclaration
2020/11/24 03:48:00 EnterEveryRule
2020/11/24 03:48:00 EnterDeclarationSpecifiers
2020/11/24 03:48:00 EnterEveryRule
2020/11/24 03:48:00 EnterDeclarationSpecifier
2020/11/24 03:48:00 EnterEveryRule
2020/11/24 03:48:00 EnterTypeSpecifier
2020/11/24 03:48:00 VisitTerminal
2020/11/24 03:48:00 ExitTypeSpecifier void
2020/11/24 03:48:00 ExitEveryRule
2020/11/24 03:48:00 ExitDeclarationSpecifier void
2020/11/24 03:48:00 ExitEveryRule
2020/11/24 03:48:00 ExitDeclarationSpecifiers void
2020/11/24 03:48:00 ExitEveryRule
2020/11/24 03:48:00 EnterEveryRule
2020/11/24 03:48:00 EnterInitDeclaratorList
2020/11/24 03:48:00 EnterEveryRule
2020/11/24 03:48:00 EnterInitDeclarator
2020/11/24 03:48:00 EnterEveryRule
2020/11/24 03:48:00 EnterDeclarator
2020/11/24 03:48:00 EnterEveryRule
2020/11/24 03:48:00 EnterDirectDeclarator
2020/11/24 03:48:00 VisitTerminal
2020/11/24 03:48:00 ExitDirectDeclarator main
2020/11/24 03:48:00 ExitEveryRule
2020/11/24 03:48:00 EnterEveryRule
2020/11/24 03:48:00 EnterDirectDeclarator
2020/11/24 03:48:00 VisitTerminal
2020/11/24 03:48:00 VisitTerminal
2020/11/24 03:48:00 ExitDirectDeclarator main()
2020/11/24 03:48:00 ExitEveryRule
2020/11/24 03:48:00 ExitDeclarator main()
2020/11/24 03:48:00 ExitEveryRule
2020/11/24 03:48:00 ExitInitDeclarator main()
2020/11/24 03:48:00 ExitEveryRule
2020/11/24 03:48:00 ExitInitDeclaratorList main()
2020/11/24 03:48:00 ExitEveryRule
line 5:1 mismatched input '{' expecting ';'
2020/11/24 03:48:00 VisitErrorNode
2020/11/24 03:48:00 VisitErrorNode
2020/11/24 03:48:00 VisitErrorNode
2020/11/24 03:48:00 VisitErrorNode
2020/11/24 03:48:00 VisitErrorNode
2020/11/24 03:48:00 VisitErrorNode
2020/11/24 03:48:00 VisitErrorNode
2020/11/24 03:48:00 VisitErrorNode
2020/11/24 03:48:00 VisitErrorNode
2020/11/24 03:48:00 VisitErrorNode
2020/11/24 03:48:00 VisitErrorNode
2020/11/24 03:48:00 VisitErrorNode
2020/11/24 03:48:00 VisitErrorNode
2020/11/24 03:48:00 VisitErrorNode
2020/11/24 03:48:00 VisitErrorNode
2020/11/24 03:48:00 VisitErrorNode
2020/11/24 03:48:00 VisitErrorNode
2020/11/24 03:48:00 VisitErrorNode
2020/11/24 03:48:00 VisitErrorNode
2020/11/24 03:48:00 VisitErrorNode
2020/11/24 03:48:00 VisitErrorNode
2020/11/24 03:48:00 ExitDeclaration voidmain(){inta=3;intb=4;printf("a+b = %d\n",a+b);}
2020/11/24 03:48:00 ExitEveryRule
(declaration (declarationSpecifiers (declarationSpecifier (typeSpecifier void))) (initDeclaratorList (initDeclarator (declarator (directDeclarator (directDeclarator main) ( ))))) { int a = 3 ; int b = 4 ; printf ( "a+b = %d\n" , a + b ) ; })
PASS
ok      _/home/b175949/parsar   0.015s
b175949@trumpet:~/parsar$ 
```
