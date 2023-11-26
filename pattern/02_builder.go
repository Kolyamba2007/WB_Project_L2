package pattern

/*
Можно применить:

-Когда код должен создавать разные представления какого-то объекта.
Например, деревянные и железобетонные дома.
-Когда нужно собирать сложные составные объекты, например, деревья Компоновщика.

Плюсы:
-Позволяет создавать продукты пошагово.
-Позволяет использовать один и тот же код для создания различных продуктов.
-Изолирует сложный код сборки продукта от его основной бизнес-логики.

Минусы:
-Усложняет код программы из-за введения дополнительных классов.
-Клиент будет привязан к конкретным классам строителей,
так как в интерфейсе директора может не быть метода получения результата.

Ниже представлен пример строителя, делающего коктейли и записыващего их рецепты.
Клиент может сам "собрать" коктейль, либо обратиться к директору и получить "заготовку".
*/

type IBuilder interface {
	SetMilkCount(int)
	SetSauceType(string)
	SetSugar(bool)
	GetResult() Product
}

func GetBuilder(builderType string) IBuilder {
	if builderType == "cocktail" {
		return newCocktailBuilder()
	}

	if builderType == "recipe" {
		return newRecipeBuilder()
	}
	return nil
}

type Product struct {
	MilkCount int
	SauceType string
	Sugar     bool
}

type CocktailBuilder struct {
	MilkCount int
	SauceType string
	Sugar     bool
}

func newCocktailBuilder() *CocktailBuilder {
	return &CocktailBuilder{}
}

func (b *CocktailBuilder) SetMilkCount(count int) {
	b.MilkCount = count
}

func (b *CocktailBuilder) SetSauceType(sauceType string) {
	b.SauceType = sauceType
}

func (b *CocktailBuilder) SetSugar(ok bool) {
	b.Sugar = ok
}

func (b *CocktailBuilder) GetResult() Product {
	return Product{
		MilkCount: b.MilkCount,
		SauceType: b.SauceType,
		Sugar:     b.Sugar,
	}
}

type RecipeBuilder struct {
	MilkCount int
	SauceType string
	Sugar     bool
}

func newRecipeBuilder() *RecipeBuilder {
	return &RecipeBuilder{}
}

func (b *RecipeBuilder) SetMilkCount(count int) {
	b.MilkCount = count
}

func (b *RecipeBuilder) SetSauceType(sauceType string) {
	b.SauceType = sauceType
}

func (b *RecipeBuilder) SetSugar(ok bool) {
	b.Sugar = ok
}

func (b *RecipeBuilder) GetResult() Product {
	return Product{
		MilkCount: b.MilkCount,
		SauceType: b.SauceType,
		Sugar:     b.Sugar,
	}
}

type Director struct {
	builder IBuilder
}

func NewDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) SetBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) ConstructCaramelLatte() Product {
	d.builder.SetMilkCount(1)
	d.builder.SetSauceType("caramel")
	d.builder.SetSugar(true)
	return d.builder.GetResult()
}
