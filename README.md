# ANTLR

### プログラム
```
	int summation(int a, int b) {
         int sum;
         sum = a+b;
         return sum;
	}
```

### コメント
- 型定義

(declarationSpecifiers (declarationSpecifier (typeSpecifier 型名)))

- 引数名？

(declarator (directDeclarator 変数名))

- 変数定義

(declarationSpecifiers (declarationSpecifier (typeSpecifier 型名)))

(initDeclarationList (initDeclarator (declarator (directDeclarator 変数名))))

- 定義済みの変数を使用するとき？

(unaryExpression (postfixExpression(primaryExpression 変数名)))

### 実行結果
```
b175949@trumpet:~/parsar$ go test
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterCompilationUnit
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterTranslationUnit
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterExternalDeclaration
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterFunctionDefinition
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterDeclarationSpecifiers
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterDeclarationSpecifier
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterTypeSpecifier
2020/11/26 08:05:13 VisitTerminal
2020/11/26 08:05:13 ExitTypeSpecifier int
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitDeclarationSpecifier int
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitDeclarationSpecifiers int
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterDeclarator
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterDirectDeclarator
2020/11/26 08:05:13 VisitTerminal
2020/11/26 08:05:13 ExitDirectDeclarator summation
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterDirectDeclarator
2020/11/26 08:05:13 VisitTerminal
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterParameterTypeList
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterParameterList
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterParameterDeclaration
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterDeclarationSpecifiers
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterDeclarationSpecifier
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterTypeSpecifier
2020/11/26 08:05:13 VisitTerminal
2020/11/26 08:05:13 ExitTypeSpecifier int
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitDeclarationSpecifier int
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitDeclarationSpecifiers int
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterDeclarator
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterDirectDeclarator
2020/11/26 08:05:13 VisitTerminal
2020/11/26 08:05:13 ExitDirectDeclarator a
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitDeclarator a
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitParameterDeclaration
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitParameterList
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterParameterList
2020/11/26 08:05:13 VisitTerminal
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterParameterDeclaration
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterDeclarationSpecifiers
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterDeclarationSpecifier
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterTypeSpecifier
2020/11/26 08:05:13 VisitTerminal
2020/11/26 08:05:13 ExitTypeSpecifier int
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitDeclarationSpecifier int
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitDeclarationSpecifiers int
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterDeclarator
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterDirectDeclarator
2020/11/26 08:05:13 VisitTerminal
2020/11/26 08:05:13 ExitDirectDeclarator b
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitDeclarator b
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitParameterDeclaration
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitParameterList
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitParameterTypeList
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 VisitTerminal
2020/11/26 08:05:13 ExitDirectDeclarator summation(inta,intb)
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitDeclarator summation(inta,intb)
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterCompoundStatement
2020/11/26 08:05:13 VisitTerminal
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterBlockItemList
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterBlockItem
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterDeclaration
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterDeclarationSpecifiers
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterDeclarationSpecifier
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterTypeSpecifier
2020/11/26 08:05:13 VisitTerminal
2020/11/26 08:05:13 ExitTypeSpecifier int
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitDeclarationSpecifier int
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitDeclarationSpecifiers int
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterInitDeclaratorList
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterInitDeclarator
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterDeclarator
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterDirectDeclarator
2020/11/26 08:05:13 VisitTerminal
2020/11/26 08:05:13 ExitDirectDeclarator sum
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitDeclarator sum
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitInitDeclarator sum
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitInitDeclaratorList sum
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 VisitTerminal
2020/11/26 08:05:13 ExitDeclaration intsum;
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitBlockItem
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitBlockItemList
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterBlockItemList
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterBlockItem
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterStatement
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterExpressionStatement
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterExpression
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterAssignmentExpression
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterUnaryExpression
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterPostfixExpression
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterPrimaryExpression
2020/11/26 08:05:13 VisitTerminal
2020/11/26 08:05:13 ExitPrimaryExpression
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitPostfixExpression
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitUnaryExpression
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterAssignmentOperator
2020/11/26 08:05:13 VisitTerminal
2020/11/26 08:05:13 ExitAssignmentOperator
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterAssignmentExpression
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterConditionalExpression
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterLogicalOrExpression
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterLogicalAndExpression
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterInclusiveOrExpression
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterExclusiveOrExpression
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterAndExpression
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterEqualityExpression
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterRelationalExpression
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterShiftExpression
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterAdditiveExpression
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterMultiplicativeExpression
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterCastExpression
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterUnaryExpression
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterPostfixExpression
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterPrimaryExpression
2020/11/26 08:05:13 VisitTerminal
2020/11/26 08:05:13 ExitPrimaryExpression
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitPostfixExpression
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitUnaryExpression
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitCastExpression
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitMultiplicativeExpression
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitAdditiveExpression
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterAdditiveExpression
2020/11/26 08:05:13 VisitTerminal
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterMultiplicativeExpression
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterCastExpression
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterUnaryExpression
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterPostfixExpression
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterPrimaryExpression
2020/11/26 08:05:13 VisitTerminal
2020/11/26 08:05:13 ExitPrimaryExpression
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitPostfixExpression
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitUnaryExpression
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitCastExpression
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitMultiplicativeExpression
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitAdditiveExpression
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitShiftExpression
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitRelationalExpression
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitEqualityExpression
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitAndExpression
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitExclusiveOrExpression
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitInclusiveOrExpression
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitLogicalAndExpression
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitLogicalOrExpression
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitConditionalExpression
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitAssignmentExpression
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitAssignmentExpression
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitExpression
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 VisitTerminal
2020/11/26 08:05:13 ExitExpressionStatement
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitStatement
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitBlockItem
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitBlockItemList
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterBlockItemList
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterBlockItem
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterStatement
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterJumpStatement
2020/11/26 08:05:13 VisitTerminal
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterExpression
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterAssignmentExpression
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterConditionalExpression
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterLogicalOrExpression
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterLogicalAndExpression
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterInclusiveOrExpression
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterExclusiveOrExpression
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterAndExpression
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterEqualityExpression
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterRelationalExpression
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterShiftExpression
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterAdditiveExpression
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterMultiplicativeExpression
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterCastExpression
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterUnaryExpression
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterPostfixExpression
2020/11/26 08:05:13 EnterEveryRule
2020/11/26 08:05:13 EnterPrimaryExpression
2020/11/26 08:05:13 VisitTerminal
2020/11/26 08:05:13 ExitPrimaryExpression
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitPostfixExpression
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitUnaryExpression
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitCastExpression
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitMultiplicativeExpression
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitAdditiveExpression
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitShiftExpression
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitRelationalExpression
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitEqualityExpression
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitAndExpression
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitExclusiveOrExpression
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitInclusiveOrExpression
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitLogicalAndExpression
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitLogicalOrExpression
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitConditionalExpression
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitAssignmentExpression
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitExpression
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 VisitTerminal
2020/11/26 08:05:13 ExitJumpStatement
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitStatement
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitBlockItem
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitBlockItemList
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 VisitTerminal
2020/11/26 08:05:13 ExitCompoundStatement
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitFunctionDefinition
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitExternalDeclaration
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 ExitTranslationUnit
2020/11/26 08:05:13 ExitEveryRule
2020/11/26 08:05:13 VisitTerminal
2020/11/26 08:05:13 ExitCompilationUnit
2020/11/26 08:05:13 ExitEveryRule
(compilationUnit (translationUnit (externalDeclaration (functionDefinition (declarationSpecifiers (declarationSpecifier (typeSpecifier int))) (declarator (directDeclarator (directDeclarator summation) ( (parameterTypeList (parameterList (parameterList (parameterDeclaration (declarationSpecifiers (declarationSpecifier (typeSpecifier int))) (declarator (directDeclarator a)))) , (parameterDeclaration (declarationSpecifiers (declarationSpecifier (typeSpecifier int))) (declarator (directDeclarator b))))) ))) (compoundStatement { (blockItemList (blockItemList (blockItemList (blockItem (declaration (declarationSpecifiers (declarationSpecifier (typeSpecifier int))) (initDeclaratorList (initDeclarator (declarator (directDeclarator sum)))) ;))) (blockItem (statement (expressionStatement (expression (assignmentExpression (unaryExpression (postfixExpression (primaryExpression sum))) (assignmentOperator =) (assignmentExpression (conditionalExpression (logicalOrExpression (logicalAndExpression (inclusiveOrExpression (exclusiveOrExpression (andExpression (equalityExpression (relationalExpression (shiftExpression (additiveExpression (additiveExpression (multiplicativeExpression (castExpression (unaryExpression (postfixExpression (primaryExpression a)))))) + (multiplicativeExpression (castExpression (unaryExpression (postfixExpression (primaryExpression b)))))))))))))))))) ;)))) (blockItem (statement (jumpStatement return (expression (assignmentExpression (conditionalExpression (logicalOrExpression (logicalAndExpression (inclusiveOrExpression (exclusiveOrExpression (andExpression (equalityExpression (relationalExpression (shiftExpression (additiveExpression (multiplicativeExpression (castExpression (unaryExpression (postfixExpression (primaryExpression sum))))))))))))))))) ;)))) })))) <EOF>)
PASS
ok      _/home/b175949/parsar   0.063s
b175949@trumpet:~/parsar$ 
```
