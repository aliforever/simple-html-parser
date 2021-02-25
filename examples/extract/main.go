package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"

	simple_html_parser "github.com/aliforever/simple-html-parser"

	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	f, _ := ioutil.ReadFile("data.html")
	html := string(f)
	// html := `<div class="SwitchableGroup js-content-key_facts is-visible" data-id="key_facts"><button class="ToggleButton js-contentButtons js-content-key_facts mdc-layout-grid" data-id="key_facts"><div class="ToggleButtonInner"><span class="ToggleButtonTitle">Key information</span><i class="ToggleButtonIcon lnr-chevron-down"></i></div></button><article class="ModulesContainer"><div data-module="Study:study_key_facts" class="Module StudyPortals_Shared_Modules_BNW_DetailedInformation_KeyFacts_KeyFacts"> <section> <div id="StudyKeyFacts"> <h2 class="KeyFactsTitle">Key information</h2> <article class="FactItem"> <h3 class="FactItemTitle">Duration</h3> <ul class="FactList DurationList">  <li class="FactListItem"> <div class="FactListTitle js-durationFact"> Full-time <i class="FactListTitleArrow lnr-arrow-right"></i> </div> <ul class="FactListSubList">  <li class="FactListSubListItem"> <span class="Duration" data-rewrite="true" data-period="months" data-duration="18" data-days-single="day" data-days-multiple="days" data-weeks-single="week" data-weeks-multiple="weeks" data-months-single="month" data-months-multiple="months" data-years-single="year" data-years-multiple="years" title=" 18&nbsp; months "> 18&nbsp;months </span> </li>  </ul> </li>  </ul> </article> <article class="FactItem"> <h3 class="FactItemTitle">Start dates &amp; application deadlines</h3>  <div>You can apply for and start this programme anytime.</div>  </article>  <article class="FactItem LanguageFact js-languageFact"> <h3 class="FactItemTitle">Language</h3> <div class="Languages"> English </div> <button class="NavigatorButton TextOnly languageRequirementsLink js-scrollToRequirements"> Check English test score requirements </button>  </article>  <article class="FactItem"> <h3 class="FactItemTitle">Credits</h3>  <div>60 <span title="European Credit Transfer and Accumulation System">ECTS</span></div>  </article> <article class="FactItem"> <h3 class="FactItemTitle">Delivered</h3> <div>  On Campus </div>  </article>  <article class="FactItem Disciplines"> <h3 class="FactItemTitle">Disciplines</h3>  <a class="TextOnly" href="/disciplines/45/business-administration.html" title="Business Administration" target="_blank"> Business Administration </a>  <a class="TextOnly" href="/disciplines/234/business-intelligence-analytics.html" title="Business Intelligence &amp; Analytics" target="_blank"> Business Intelligence &amp; Analytics </a>  <a class="TextOnly" href="/disciplines/282/data-science-big-data.html" title="Data Science &amp; Big Data" target="_blank"> Data Science &amp; Big Data </a>  <a class="TextOnly LandingPageLink" href="/study-options/268828954/data-science-big-data-spain.html" title="View 17 other Masters in Data Science &amp; Big Data in Spain" target="_blank"> View 17 other Masters in Data Science &amp; Big Data in Spain </a>  </article> </div> <article class="ProgrammeWebsiteContainer"> <h3 class="Title">Explore more key information</h3> <a href="https://sl.prtl.co/track/click/?target=https%3A%2F%2Feneb.com%2Fspecial-offer-bdbi%2F&amp;facts=eyJsIjp7InQiOiJzdHVkeSIsImkiOiIzMDUyMjYiLCJkIjoiQmlnIERhdGEgYW5kIEJ1c2luZXNzIEludGVsbGlnZW5jZSIsImwiOiJwcmVtaXVtIn0sInMiOm51bGwsInUiOnsiaSI6IjE3Mi4zMC45LjE1OCIsInMiOiJlZjA3YTdjMi1hZmYxLTRkNTEtOWI4MS1mMWU5Mzg4OGJhM2UiLCJsIjoiZW4tR0IiLCJjIjoiaXIifSwiZiI6W3siYSI6ImNsaWMiLCJ0IjoibGluayIsImkiOiIyMTYyNjk4IiwiZCI6Imh0dHBzOi8vZW5lYi5jb20vc3BlY2lhbC1vZmZlci1iZGJpLyIsImwiOiJyZXZlbnVlIiwiZXgiOnsicHQiOiJwIn19XX0%3D&amp;taps=null&amp;duid=f4e9fe4e-5aac-4610-9432-b21ec7d51009&amp;sid=99d0a55c-0d30-45b0-bd61-65ad9fd0667a&amp;uid=eu-west-1%3A32b8f50c-2df4-4141-b327-58288e7e4c88" title="Visit official programme website" class="StudyLink TrackingExternalLink ProgrammeWebsiteLink" target="mp_external_2162698" rel="noopener" data-ga-tracking="{&quot;category&quot;:&quot;Study&quot;,&quot;action&quot;:&quot;Premium Click&quot;,&quot;label&quot;:&quot;305226&quot;,&quot;value&quot;:&quot;key_facts_tab&quot;}" data-action="clic" data-type="link" data-id="2162698" data-description="https://eneb.com/special-offer-bdbi/" data-label="revenue" data-registration-lockable="false" data-extra="{&quot;pt&quot;:&quot;p&quot;}" data-ga-parsed="true">Visit official programme website</a> </article>  </section>  </div></article></div>`
	p := simple_html_parser.NewParser(html)
	body, err := p.ExtractTag(`<h1 class="StudyTitle">`)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(body)
}
