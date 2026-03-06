# Feasibility Study: PulseAI Insights

## Table of Contents
1. [Data Source Viability](#1-data-source-viability)
2. [AI Processing Costs & Token Economics](#2-ai-processing-costs--token-economics)
3. [Monetization Automation Feasibility](#3-monetization-automation-feasibility)
4. [Platform & API Restrictions](#4-platform--api-restrictions)
5. [Anonymity & Compliance Verification](#5-anonymity--compliance-verification)

---

## 1. Data Source Viability
- [x] Research 5-10 high-authority healthcare/AI RSS feeds or APIs.
- [x] Verify if these sources allow automated scraping/repurposing (Terms of Service check).
- [ ] Test data extraction quality from at least 3 sources.

## 2. AI Processing Costs & Token Economics
- [x] Calculate average token count for a typical healthcare news summary.
- [x] Compare costs between Gemini Flash, GPT-4o-mini, and Claude Haiku for daily runs.
- [x] Estimate monthly 'Brain' cost based on 10 articles/day.

## 3. Monetization Automation Feasibility
- [x] Research Amazon Associates API requirements (minimum sales for API access).
- [x] Investigate automated AdSense approval requirements for AI-generated content.
- [ ] Identify 3 niche healthcare affiliate programs with public APIs or deep-linking.

## 4. Platform & API Restrictions
- [x] Check X (Twitter) API v2 free tier limits for automated posting.
- [x] Verify LinkedIn API access requirements for automated company page updates.
- [x] Confirm GitHub Actions free tier limits for daily cron jobs.

## 5. Anonymity & Compliance Verification
- [x] Research domain registrars with the best WHOIS privacy and crypto/anonymous payment options.
- [x] Verify legal implications of AI-generated medical news summaries (disclaimers needed).

### Findings: Data Source Viability
- **PubMed:** Highly viable. Public domain data. RSS feeds available for specific MeSH terms (e.g., "Artificial Intelligence").
- **BMJ Digital Health & AI:** Viable via Open Access (CC BY) content. Requires careful attribution. Automated scraping of the *site* is restricted, but RSS snippets are fair game.
- **ScienceDaily / Medgadget:** Viable for short summaries based on RSS snippets. Full article scraping is restricted.
- **MIT Tech Review:** High quality but strict. Best used as a 'Scout' source to find topics, then use PubMed for the actual content generation.

### Findings: AI Processing Costs
- **Model Choice:** Gemini 2.0 Flash is the recommended 'Brain' due to its low cost ($0.10/1M input) and high speed.
- **Daily Token Estimate:** ~6,000 input tokens, ~1,500 output tokens.
- **Daily Cost:** ~$0.0012.
- **Monthly Cost:** ~$0.036 (less than 5 cents).
- **Conclusion:** The 'Brain' cost is negligible. The primary costs will be domain registration (~$10-15/year).

### Findings: Monetization Automation
- **Amazon Associates:** PA API access requires 3 sales in 180 days. 
    - *Strategy:* Use AI to generate 'Deep Links' (standard affiliate URLs) initially. Switch to API once the sales threshold is met.
- **Google AdSense:** Feasible. AI content is allowed if it provides unique value (summarization/curation counts). Requires a 'clean' site with standard pages (About, Privacy, Contact).
- **Niche Affiliates:** Programs like HealthAffiliate.com or specialized AI tool affiliates (e.g., Jasper, Copy.ai) often have simpler deep-linking that AI can easily automate.

### Findings: Platform & API Restrictions
- **X (Twitter):** Free tier is viable for 1-2 posts/day (500 posts/month limit). Write-only access is sufficient for our needs.
- **GitHub Actions:** Highly viable. 2,000 free minutes/month is more than enough for a daily 5-minute curation script.
- **LinkedIn:** (Pending) Likely requires 'Marketing Developer Platform' access for automated company page posting. Fallback: Use a tool like Buffer or Make.com (free tiers) if direct API access is too complex.

### Findings: Anonymity & Compliance
- **Registrars:** 
    - *Njalla:* Best for anonymity (no personal info, crypto accepted).
    - *OrangeWebsite:* Good alternative, Iceland-based, privacy-focused.
    - *Namecheap:* Easiest mainstream option with free WHOIS privacy.
- **Compliance:** 
    - *AI Disclosure:* Must clearly state content is AI-curated/summarized.
    - *Medical Disclaimer:* Mandatory "Not medical advice. Consult a professional." footer on every page.
    - *Copyright:* Summarization (transformative use) with attribution is generally safe under Fair Use, but we must link back to original sources.

## Final Feasibility Verdict: ✅ HIGHLY FEASIBLE
The project has extremely low overhead (~$0.04/mo for AI, ~$15/yr for domain) and high automation potential. The primary hurdle is the initial 3 sales for Amazon API access, which can be bypassed using deep links.
