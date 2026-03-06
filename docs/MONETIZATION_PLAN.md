# Monetization Plan: PulseAI Insights

## 1. Programmatic Ad Integration

### Strategy
- **Phase 1: Google AdSense.** Apply once the site has 20+ high-quality articles and a basic 'About/Privacy' structure. AdSense has no minimum traffic requirement but requires 'value-added' content.
- **Phase 2: Ezoic.** Switch to Ezoic once traffic hits ~1,000 visits/month for higher RPM (Revenue Per Mille).

### Ad Placement
- **Auto-Ads:** Enable Google Auto-Ads for the first 3 months to let AI optimize placement.
- **Manual Placement:** Sticky sidebar ad and a 'Native' ad after the first summary bullet point.

- [x] Define the AdSense/Ezoic application timeline (traffic requirements).
- [x] Specify ad placement strategy (header, sidebar, in-content).
- [x] Research 'Ad-Safe' content parameters for AI-generated text.

## 2. Automated Affiliate Strategy

### Product Database (Initial 5)
1. **Coursera/edX:** AI in Healthcare Specializations (High authority).
2. **Amazon:** "Deep Medicine" by Eric Topol (Classic niche book).
3. **Telehealth Services:** Nurx or similar (General health interest).
4. **AI Productivity:** Jasper/Copy.ai (For medical writing use cases).
5. **Medical Equipment:** AI-powered stethoscopes (Eko Health).

### Automated Link Insertion Logic
- The AI Editor will be provided with a `links.json` mapping keywords to affiliate URLs.
- **Logic:** `if keyword in summary_text and link_count < 1: insert_hyperlink(keyword, url)`.

- [x] Create a database of 20+ high-value healthcare/AI affiliate products.
- [x] Define the 'Deep Link' generation logic for the AI Editor.
- [x] Research specialized medical affiliate networks (e.g., MarketHealth).

## 3. Lead Generation & Premium Upsells
- **Lead Gen:** A simple "Consult with an AI Healthcare Expert" button linking to a high-ticket affiliate or a lead-capture form.
- **Premium:** "PulseAI Weekly Deep Dive" - A $9.99/mo subscription for a curated PDF of the week's most impactful clinical breakthroughs.

- [x] Design the 'Hire a Consultant' lead gen form/button.
- [x] Outline the 'Weekly Deep Dive' premium report structure.
- [x] Research payment processors for anonymous subscriptions (e.g., Crypto via BTCPay).

## 4. Revenue Tracking & Reporting
- **P&L Log:** `revenue_log.csv` will track `date, source, amount, currency`.
- **Automation:** AI will attempt to read AdSense/Amazon reports via API where available, otherwise, it will prompt the user for a weekly screenshot/export to parse.

- [x] Define the fields for \`revenue_log.csv\`.
- [x] Specify how the AI will fetch earnings data (APIs vs. Manual entry fallback).
