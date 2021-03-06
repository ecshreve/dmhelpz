// Code generated by github.com/Yamashou/gqlgenc, DO NOT EDIT.

package gen

import (
	"context"
	"net/http"

	"github.com/Yamashou/gqlgenc/client"
)

type Client struct {
	Client *client.Client
}

func NewClient(cli *http.Client, baseURL string, options ...client.HTTPRequestOption) *Client {
	return &Client{Client: client.NewClient(cli, baseURL, options...)}
}

type Query struct {
	AbilityScore        *AbilityScore        "json:\"abilityScore\" graphql:\"abilityScore\""
	AbilityScores       []*AbilityScore      "json:\"abilityScores\" graphql:\"abilityScores\""
	Alignment           *Alignment           "json:\"alignment\" graphql:\"alignment\""
	Alignments          []*Alignment         "json:\"alignments\" graphql:\"alignments\""
	Background          *Background          "json:\"background\" graphql:\"background\""
	Backgrounds         []*Background        "json:\"backgrounds\" graphql:\"backgrounds\""
	Condition           *Condition           "json:\"condition\" graphql:\"condition\""
	Conditions          []*Condition         "json:\"conditions\" graphql:\"conditions\""
	Class               *Class               "json:\"class\" graphql:\"class\""
	Classes             []*Class             "json:\"classes\" graphql:\"classes\""
	DamageType          *DamageType          "json:\"damageType\" graphql:\"damageType\""
	DamageTypes         []*DamageType        "json:\"damageTypes\" graphql:\"damageTypes\""
	Equipment           *Equipment           "json:\"equipment\" graphql:\"equipment\""
	Equipments          []*Equipment         "json:\"equipments\" graphql:\"equipments\""
	EquipmentCategory   *EquipmentCategory   "json:\"equipmentCategory\" graphql:\"equipmentCategory\""
	EquipmentCategories []*EquipmentCategory "json:\"equipmentCategories\" graphql:\"equipmentCategories\""
	Feat                *Feat                "json:\"feat\" graphql:\"feat\""
	Feats               []*Feat              "json:\"feats\" graphql:\"feats\""
	Feature             *Feature             "json:\"feature\" graphql:\"feature\""
	Features            []*Feature           "json:\"features\" graphql:\"features\""
	Language            *Language            "json:\"language\" graphql:\"language\""
	Languages           []*Language          "json:\"languages\" graphql:\"languages\""
	Level               *Level               "json:\"level\" graphql:\"level\""
	Levels              []*Level             "json:\"levels\" graphql:\"levels\""
	MagicItem           *MagicItem           "json:\"magicItem\" graphql:\"magicItem\""
	MagicItems          []*MagicItem         "json:\"magicItems\" graphql:\"magicItems\""
	MagicSchool         *MagicSchool         "json:\"magicSchool\" graphql:\"magicSchool\""
	MagicSchools        []*MagicSchool       "json:\"magicSchools\" graphql:\"magicSchools\""
	Monster             *Monster             "json:\"monster\" graphql:\"monster\""
	Monsters            []*Monster           "json:\"monsters\" graphql:\"monsters\""
	Proficiency         *Proficiency         "json:\"proficiency\" graphql:\"proficiency\""
	Proficiencies       []*Proficiency       "json:\"proficiencies\" graphql:\"proficiencies\""
	Race                *Race                "json:\"race\" graphql:\"race\""
	Races               []*Race              "json:\"races\" graphql:\"races\""
	Rule                *Rule                "json:\"rule\" graphql:\"rule\""
	Rules               []*Rule              "json:\"rules\" graphql:\"rules\""
	RuleSection         *RuleSection         "json:\"ruleSection\" graphql:\"ruleSection\""
	RuleSections        []*RuleSection       "json:\"ruleSections\" graphql:\"ruleSections\""
	Skill               *Skill               "json:\"skill\" graphql:\"skill\""
	Skills              []*Skill             "json:\"skills\" graphql:\"skills\""
	Spell               *Spell               "json:\"spell\" graphql:\"spell\""
	Spells              []*Spell             "json:\"spells\" graphql:\"spells\""
	Subclass            *Subclass            "json:\"subclass\" graphql:\"subclass\""
	Subclasses          []*Subclass          "json:\"subclasses\" graphql:\"subclasses\""
	Subrace             *Subrace             "json:\"subrace\" graphql:\"subrace\""
	Subraces            []*Subrace           "json:\"subraces\" graphql:\"subraces\""
	Trait               *Trait               "json:\"trait\" graphql:\"trait\""
	Traits              []*Trait             "json:\"traits\" graphql:\"traits\""
	WeaponProperty      *WeaponProperty      "json:\"weaponProperty\" graphql:\"weaponProperty\""
	WeaponProperties    []*WeaponProperty    "json:\"weaponProperties\" graphql:\"weaponProperties\""
}
type GetMonster struct {
	Monster *struct {
		Name  *string "json:\"name\" graphql:\"name\""
		Speed *struct {
			Burrow *string "json:\"burrow\" graphql:\"burrow\""
			Walk   *string "json:\"walk\" graphql:\"walk\""
		} "json:\"speed\" graphql:\"speed\""
	} "json:\"monster\" graphql:\"monster\""
}

const GetMonsterDocument = `query GetMonster ($name: String!) {
	monster(filter: {name:$name}) {
		name
		speed {
			burrow
			walk
		}
	}
}
`

func (c *Client) GetMonster(ctx context.Context, name string, httpRequestOptions ...client.HTTPRequestOption) (*GetMonster, error) {
	vars := map[string]interface{}{
		"name": name,
	}

	var res GetMonster
	if err := c.Client.Post(ctx, "GetMonster", GetMonsterDocument, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}
