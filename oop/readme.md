# 面向对象编程-继承
    
    type Goods struct{
        Name string
        Price int
    }
    
    type Book struct{
        Goods
        Writer string
    }
    
# 项目开发流程说明
>家庭记账软件

- 记录家庭的收入、支出
- 打印收支明细表
- 退出

思路分析：
**把记账软件的功能，封装到一个结构体中，然后调用该结构体的方法，来实现记账，显示明细。结构体的名字FamilyAccount**
>在通过在main方法中，创建一个结构体FamilyAccount实例，实现记账即可。