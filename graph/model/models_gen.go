// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type EmploymentRate struct {
	Meta *Meta                 `json:"Meta"`
	Data []*EmploymentRateData `json:"Data"`
}

type EmploymentRateData struct {
	Timeperiod       *string `json:"timeperiod"`
	Maleemployment   *string `json:"maleemployment"`
	Femaleemployment *string `json:"femaleemployment"`
	Maleaged1617     *string `json:"maleaged1617"`
	Maleaged1824     *string `json:"maleaged1824"`
	Maleaged2534     *string `json:"maleaged2534"`
	Maleaged3549     *string `json:"maleaged3549"`
	Maleaged5064     *string `json:"maleaged5064"`
	Maleaged65       *string `json:"maleaged65"`
	Femaleaged1617   *string `json:"femaleaged1617"`
	Femaleaged1824   *string `json:"femaleaged1824"`
	Femaleaged2534   *string `json:"femaleaged2534"`
	Femaleaged3549   *string `json:"femaleaged3549"`
	Femaleaged5064   *string `json:"femaleaged5064"`
	Femaleaged65     *string `json:"femaleaged65"`
	Allaged1617      *string `json:"allaged1617"`
	Allaged1824      *string `json:"allaged1824"`
	Allaged2534      *string `json:"allaged2534"`
	Allaged3549      *string `json:"allaged3549"`
	Allaged5064      *string `json:"allaged5064"`
}

type GenderIdentity struct {
	Location           *string `json:"location"`
	Allothergenders    *string `json:"allothergenders"`
	Differentfrombirth *string `json:"differentfrombirth"`
	Sameasbirth        *string `json:"sameasbirth"`
	Nonbinary          *string `json:"nonbinary"`
	Transman           *string `json:"transman"`
	Transwoman         *string `json:"transwoman"`
	Unanswered         *string `json:"unanswered"`
	Total              *string `json:"total"`
}

type Meta struct {
	Source      *string `json:"Source"`
	LastUpdated *string `json:"LastUpdated"`
	NextRelease *string `json:"NextRelease"`
}

type Population struct {
	Timeperiod *string `json:"timeperiod"`
	Ukpop      *string `json:"ukpop"`
	Engpop     *string `json:"engpop"`
	Wapop      *string `json:"wapop"`
	Nipop      *string `json:"nipop"`
	Scopop     *string `json:"scopop"`
}
