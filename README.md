# TextFormer
## Ветвь (PKG branch) для подключения в качестве PKG   
Программа для добавления в текст новых блоков   
``` go
type TextString string      // простой текстовый блок   
// Структура текстовых блоков:   
> type TextTemplate struct  // текстовый блок с разметкой   
> {  
> {  
>	  Top         TextString // верхняя часть   
>	  TopSplit    TextString // разделитель   
>	  Middle      TextString // средняя часть   
>	  BottomSplit TextString // разделитель   
>	  Bottom      TextString // нижняя часть   
>}   
>   
// API для работы с текстом   
>type IText interface   
>{   
>	  Get() string // получить блок текст   
>	  Set(...string) // перезаписать блоки текст в порядке: Верхний блок, Нижний блок, Средний Блок   
>	  AddUD(...string) // Добавить в средний блок новые блоки для последующего чтения блоков сверху вниз (Как в книгах)   
>	  AddDU(...string) // Добавить в средний блок новые блоки текста   
>	                   // для последующего просмотра блоков снизу вверх (Как в чатах, блогах и форумах)   
>	  SetSplit(string, string) // Добавить разделители блоков текста (только для верхнего блока и нижнего блока)   
>	  AddTopUD(...string) // Добавить блоки текста в верхнюю часть страницы   
>	                      // для последовательного прочтения сверху вниз (Как в книгах)   
>	  AddBottomUD(...string) // Добавить блоки текста в нижнюю часть страницы   
>	                         // для последовательного прочтения сверху вниз (Как в чатах, блогах и форумах)   
>	  Clean()	// Очистка всего текстового блока (верхняя, средняя и нижняя части   
>}   
>   
```
