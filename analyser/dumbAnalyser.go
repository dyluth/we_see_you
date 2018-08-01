package analyser

var (
	// a dumb map between the searchable string, and some keywords to find in the tweet
	categories = make(map[string][]string)
)

func init() {

	categories["gay rights"] = []string{"gay", "lesbian", "LGB", "LGBT"}
	categories["smoking bans"] = []string{"smoking ban"}
	categories["marriage"] = []string{"gay marriage", "marriage", "same sex"} // allowing marriage between two people of same sex
	categories["equality and human rights"] = []string{"equality", "human rights"}
	categories["terminally ill people assistance to end their life"] = []string{"euthenasia"}
	categories["UK military forces"] = []string{"military intervention", "peacekeeping"} //use of UK military forces in combat operations overseas
	categories["investigations"] = []string{"iraq", "iraq war"}                          //investigations into the Iraq war
	categories["Trident"] = []string{"trident", "WMD", "nuclear weapons"}
	categories["EU integration"] = []string{"eu integration"}       // eu integration
	categories["EU"] = []string{"eu referendum"}                    //referendum on the UK's membership of the EU
	categories["Military Covenant"] = []string{"Military Covenant"} //strengthening the Military Covenant
	categories["right to remain for EU nationals"] = []string{"EU national", "right to remain"}
	categories["UK membership of the EU"] = []string{"EU", "brexit"}
	categories["military action against ISIL (Daesh)"] = []string{"ISIL"}
	categories["housing benefit"] = []string{"bedroom tax", "housing benefit"} // bedroom tax
	categories["welfare benefits"] = []string{"welfare", "benefits"}           //raising welfare benefits at least in line with prices
	categories["illness or disability"] = []string{"disability benefits"}      //paying higher benefits over longer periods for those unable to work due to illness or disability
	/*
		categories["financial need council tax"] = []string{}                      //making local councils responsible for helping those in financial need afford their council tax and reducing the amount spent on such support
		categories["welfare benefits"] = []string{}
		categories["guaranteed jobs for young people"] = []string{}
		categories["income tax"] = []string{}
		categories["rate of VAT"] = []string{}
		categories["alcoholic drinks"] = []string{}
		categories["taxes on plane tickets"] = []string{}
		categories["fuel for motor vehicles"] = []string{}
		categories["income over £150,000"] = []string{}
		categories["occupational pensions"] = []string{}
		categories["occupational pensions"] = []string{}
		categories["banker’s bonus tax"] = []string{}
		categories["taxes on banks"] = []string{}
		categories["mansion tax"] = []string{}
		categories["rights for shares"] = []string{}
		categories["regulation of trade union activity"] = []string{}
		categories["capital gains tax"] = []string{}
		categories["corporation tax"] = []string{}
		categories["tax avoidance"] = []string{}
		categories["incentives for companies to invest"] = []string{}
		categories["high speed rail"] = []string{}
		categories["private patients"] = []string{}
		categories["NHS"] = []string{}
		categories["smoking bans"] = []string{}
		categories["terminally ill people assistance to end their life"] = []string{}
		categories["autonomy for schools"] = []string{}
		categories["undergraduate tuition fee"] = []string{}
		categories["academy schools"] = []string{}
		categories["financial support"] = []string{}
		categories["tuition fees"] = []string{}
		categories["funding of local government"] = []string{}
		categories["equal number of electors"] = []string{}
		categories["fewer MPs"] = []string{}
		categories["proportional system"] = []string{}
		categories["wholly elected"] = []string{}
		categories["taxes on business premises"] = []string{}
		categories["campaigning by third parties"] = []string{}
		categories["fixed periods between parliamentary elections"] = []string{}
		categories["hereditary peers"] = []string{}
		categories["more powers to the Welsh Assembly"] = []string{}
		categories["more powers to the Scottish Parliament"] = []string{}
		categories["powers for local councils"] = []string{}
		categories["a veto for MPs over laws specifically impacting their part of the UK"] = []string{}
		categories["voting age"] = []string{}
		categories["stricter asylum system"] = []string{}
		categories["Police and Crime Commissioners"] = []string{}
		categories["retention of information about communications"] = []string{}
		categories["enforcement of immigration rules"] = []string{}
		categories["mass surveillance"] = []string{}
		categories["merging police and fire services"] = []string{}
		categories["prevent climate change"] = []string{}
		categories["fuel for motor vehicles"] = []string{}
		categories["forests"] = []string{}
		categories["taxes on plane tickets"] = []string{}
		categories["low carbon electricity generation"] = []string{}
		categories["culling badgers"] = []string{}
		categories["hydraulic fracturing (fracking)"] = []string{}
		categories["high speed rail"] = []string{}
		categories["bus services"] = []string{}
		categories["rail fares"] = []string{}
		categories["fuel for motor vehicles"] = []string{}
		categories["taxes on plane tickets"] = []string{}
		categories["publicly owned railway system"] = []string{}
		categories["secure tenancies for life"] = []string{}
		categories["market rent to high earners renting a council home"] = []string{}
		categories["regulation of gambling"] = []string{}
		categories["civil service redundancy payments"] = []string{}
		categories["anti-terrorism laws"] = []string{}
		categories["Royal Mail"] = []string{}
		categories["pub landlords rent-only leases"] = []string{}
		categories["legal aid"] = []string{}
		categories["courts in secret sessions"] = []string{}
		categories["register of lobbyists"] = []string{}
		categories["fees no-win no fee cases"] = []string{}
		categories["fees letting agents"] = []string{}
		categories["Conservative - Liberal Democrat Coalition Agreement"] = []string{}
	*/
}
