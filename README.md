Base on [johnfercher/maroto](https://github.com/johnfercher/maroto)

Add  function Base on [gofpdf](https://github.com/jung-kurt/gofpdf) 
1. AddUTF8Font to support chinese 
1. Add  SetProtection to support protection pdf 

站在巨人的肩膀

AddUTF8Font 方法，支持自己设定字体:比如使用 [NotoSansSC-Regular.ttf](https://github.com/jsntn/webfonts/blob/master/NotoSansSC-Regular.ttf) 就可以支持中文的展示。

SetProtection 方法，支持PDF设置密码。

maroto的布局是一行12个格子，然后向每个格子每个格子的添加数据。