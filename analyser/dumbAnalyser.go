package analyser

import (
	"strings"
)

type DumbAnalyserImpl struct {
	// a dumb map between the searchable string, and some keywords to find in the tweet
	categories map[string][]string
}

// GetCategory takes a message, and tries to deduce
func (da *DumbAnalyserImpl) GetCategory(message string) (cat string, confidence int) {

	for k, v := range da.categories {
		for _, keyword := range v {
			if strings.Contains(message, keyword) {
				return k, 90
			}
		}
	}
	return "", 0
}

// creates a new dumb analyser
func NewDumbAnalyser() DumbAnalyserImpl {
	da := DumbAnalyserImpl{}
	da.categories = make(map[string][]string)
	da.categories["gay rights"] = []string{"gay", "lesbian", "LGB", "LGBT"}
	da.categories["smoking bans"] = []string{"smoking ban"}
	da.categories["marriage"] = []string{"gay marriage", "same sex"} // allowing marriage between two people of same sex
	da.categories["equality and human rights"] = []string{"equality", "human rights"}

	da.categories["terminally ill people assistance to end their life"] = []string{"euthenasia"}
	da.categories["UK military forces"] = []string{"military intervention", "peacekeeping"} //use of UK military forces in combat operations overseas
	da.categories["investigations"] = []string{"iraq", "iraq war"}                          //investigations into the Iraq war
	da.categories["Trident"] = []string{"trident", "WMD", "nuclear weapons"}
	da.categories["EU integration"] = []string{"eu integration"}       // eu integration
	da.categories["EU"] = []string{"eu referendum", "euref"}           //referendum on the UK's membership of the EU
	da.categories["Military Covenant"] = []string{"Military Covenant"} //strengthening the Military Covenant
	da.categories["right to remain for EU nationals"] = []string{"EU national", "right to remain", "expat"}
	da.categories["UK membership of the EU"] = []string{"EU", "brexit"}
	da.categories["military action against ISIL (Daesh)"] = []string{"ISIL", "Daesh"}
	da.categories["housing benefit"] = []string{"bedroom tax", "housing benefit"} // bedroom tax
	da.categories["welfare benefits"] = []string{"welfare", "benefits"}           //raising welfare benefits at least in line with prices
	da.categories["illness or disability"] = []string{"disability benefits"}      //paying higher benefits over longer periods for those unable to work due to illness or disability
	/*
		da.categories["financial need council tax"] = []string{}                      //making local councils responsible for helping those in financial need afford their council tax and reducing the amount spent on such support
		da.categories["welfare benefits"] = []string{}
		da.categories["guaranteed jobs for young people"] = []string{}
		da.categories["income tax"] = []string{}
		da.categories["rate of VAT"] = []string{}
		da.categories["alcoholic drinks"] = []string{}
		da.categories["taxes on plane tickets"] = []string{}
		da.categories["fuel for motor vehicles"] = []string{}
		da.categories["income over £150,000"] = []string{}
		da.categories["occupational pensions"] = []string{}
		da.categories["occupational pensions"] = []string{}
		da.categories["banker’s bonus tax"] = []string{}
		da.categories["taxes on banks"] = []string{}
		da.categories["mansion tax"] = []string{}
		da.categories["rights for shares"] = []string{}
		da.categories["regulation of trade union activity"] = []string{}
		da.categories["capital gains tax"] = []string{}
		da.categories["corporation tax"] = []string{}
		da.categories["tax avoidance"] = []string{}
		da.categories["incentives for companies to invest"] = []string{}
		da.categories["high speed rail"] = []string{}
		da.categories["private patients"] = []string{}
		da.categories["NHS"] = []string{}
		da.categories["smoking bans"] = []string{}
		da.categories["terminally ill people assistance to end their life"] = []string{}
		da.categories["autonomy for schools"] = []string{}
		da.categories["undergraduate tuition fee"] = []string{}
		da.categories["academy schools"] = []string{}
		da.categories["financial support"] = []string{}
		da.categories["tuition fees"] = []string{}
		da.categories["funding of local government"] = []string{}
		da.categories["equal number of electors"] = []string{}
		da.categories["fewer MPs"] = []string{}
		da.categories["proportional system"] = []string{}
		da.categories["wholly elected"] = []string{}
		da.categories["taxes on business premises"] = []string{}
		da.categories["campaigning by third parties"] = []string{}
		da.categories["fixed periods between parliamentary elections"] = []string{}
		da.categories["hereditary peers"] = []string{}
		da.categories["more powers to the Welsh Assembly"] = []string{}
		da.categories["more powers to the Scottish Parliament"] = []string{}
		da.categories["powers for local councils"] = []string{}
		da.categories["a veto for MPs over laws specifically impacting their part of the UK"] = []string{}
		da.categories["voting age"] = []string{}
		da.categories["stricter asylum system"] = []string{}
		da.categories["Police and Crime Commissioners"] = []string{}
		da.categories["retention of information about communications"] = []string{}
		da.categories["enforcement of immigration rules"] = []string{}
		da.categories["mass surveillance"] = []string{}
		da.categories["merging police and fire services"] = []string{}
		da.categories["prevent climate change"] = []string{}
		da.categories["fuel for motor vehicles"] = []string{}
		da.categories["forests"] = []string{}
		da.categories["taxes on plane tickets"] = []string{}
		da.categories["low carbon electricity generation"] = []string{}
		da.categories["culling badgers"] = []string{}
		da.categories["hydraulic fracturing (fracking)"] = []string{}
		da.categories["high speed rail"] = []string{}
		da.categories["bus services"] = []string{}
		da.categories["rail fares"] = []string{}
		da.categories["fuel for motor vehicles"] = []string{}
		da.categories["taxes on plane tickets"] = []string{}
		da.categories["publicly owned railway system"] = []string{}
		da.categories["secure tenancies for life"] = []string{}
		da.categories["market rent to high earners renting a council home"] = []string{}
		da.categories["regulation of gambling"] = []string{}
		da.categories["civil service redundancy payments"] = []string{}
		da.categories["anti-terrorism laws"] = []string{}
		da.categories["Royal Mail"] = []string{}
		da.categories["pub landlords rent-only leases"] = []string{}
		da.categories["legal aid"] = []string{}
		da.categories["courts in secret sessions"] = []string{}
		da.categories["register of lobbyists"] = []string{}
		da.categories["fees no-win no fee cases"] = []string{}
		da.categories["fees letting agents"] = []string{}
		da.categories["Conservative - Liberal Democrat Coalition Agreement"] = []string{}
	*/
	return da
}
