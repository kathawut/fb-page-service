# Facebook App Legal Pages

This directory contains the legal documentation required for Facebook app approval.

## 📄 Pages Included

### 1. Privacy Policy (`privacy-policy.html`)
- **Purpose**: Required by Facebook for app review
- **Content**: Explains data collection, usage, and protection practices
- **URL**: `/privacy-policy`

### 2. Terms of Service (`terms-of-service.html`)
- **Purpose**: Defines user agreement and service terms
- **Content**: User responsibilities, prohibited uses, and legal compliance
- **URL**: `/terms-of-service`

## 🚀 Running the Web Server

### Start the web server:
```bash
go run cmd/web-server/main.go
```

### Or set custom port:
```bash
PORT=3000 go run cmd/web-server/main.go
```

### Access pages:
- Privacy Policy: http://localhost:3000/privacy-policy
- Terms of Service: http://localhost:3000/terms-of-service
- Root: http://localhost:3000/ (redirects to privacy policy)

## 📋 Facebook App Review Requirements

For Facebook app approval, ensure you:

1. **Update Contact Information** in both documents:
   - Replace `privacy@yourcompany.com` with your real email
   - Replace `legal@yourcompany.com` with your real email
   - Add your company address and phone number

2. **Host Publicly**: These pages must be accessible via public URLs

3. **Link from Facebook App Settings**:
   - Privacy Policy URL: `https://yourdomain.com/privacy-policy`
   - Terms of Service URL: `https://yourdomain.com/terms-of-service`

## 🔧 Customization

### Required Updates Before Production:

1. **Contact Information** (lines to update):
   - Email addresses for privacy and legal contacts
   - Company physical address
   - Phone number
   - Data Protection Officer details (if required)

2. **Legal Jurisdiction** (Terms of Service):
   - Update governing law section with your jurisdiction

3. **Company Details**:
   - Replace placeholder company references
   - Add specific business information

### Optional Customizations:

1. **Styling**: Update CSS to match your brand colors
2. **Logo**: Add company logo to the header
3. **Additional Sections**: Add any specific legal requirements for your region

## 🌐 Deployment Options

### 1. Standalone Web Server
```bash
# Production
./web-server

# Development
go run cmd/web-server/main.go
```

### 2. Add to Main API Server
You can integrate these routes into your main API server by copying the handler functions.

### 3. Static Hosting
Upload HTML files to any static hosting service (Netlify, Vercel, GitHub Pages, etc.)

## ✅ Facebook Platform Compliance

These documents are designed to comply with:
- Facebook Platform Policy
- Facebook Developer Terms
- GDPR (General Data Protection Regulation)
- CCPA (California Consumer Privacy Act)
- General data protection best practices

## 📞 Support

For questions about legal compliance or customization, consult with:
- Legal counsel familiar with data protection laws
- Facebook Developer Support
- Privacy compliance specialists

## 🔗 Useful Links

- [Facebook Platform Policy](https://developers.facebook.com/policy)
- [Facebook Developer Terms](https://developers.facebook.com/terms)
- [Facebook App Review Process](https://developers.facebook.com/docs/app-review)
- [GDPR Compliance Guide](https://gdpr.eu/)
