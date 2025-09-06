---
slug: deciding-on-what-to-do
title: Deciding on what to do
authors: [atareversei]
tags: [planning]
---
Today was a big day in shaping the direction of my IDS project. I spent most of my time thinking deeply about **how to structure the system**, what to build first, and how to balance **learning with execution**.

---

## **Key Insights**

* Building a full production-grade IDS is **complex**, involving multiple moving parts:

    * Packet sniffing
    * Flow generation and aggregation
    * Event handling
    * Backend APIs
    * Machine learning for anomaly detection
    * Dashboard for visualization
* This complexity can be overwhelming if I try to build everything at once.
* The right strategy is to **start with the core**, get it working, and expand from there.

---

## **Decision: Two-Track Approach**

I decided to **split my work into two parallel tracks**:

### **1. Core Technical Track**

The main focus will be on **building the core data pipeline**:

* **Sniffer**: Capture packets from a selected network interface and parse layers 2–4 (MAC, IP, TCP/UDP, etc.).
* **Flow Generator**: Aggregate packets into flows using the 5-tuple (SrcIP, SrcPort, DstIP, DstPort, Protocol).

I will **not worry about events, RabbitMQ, or APIs yet**.
Instead, I’ll leave placeholders in the code where those integrations will happen later.
The goal here is to **see flows form in real time** and make sure packet parsing and aggregation work correctly.

> *This is the "engine" of the IDS. Without it, the rest of the system has nothing to work with.*

---

### **2. Architecture & System Design Track**

In parallel, I’ll be **learning and planning** the broader architecture:

* Defining **data flows**, APIs, and event structures.
* Designing how services like the backend and ML pipeline will communicate.
* Writing everything down in `/docs` using **Docusaurus** to keep it organized and professional.
* Slowly building up the backend and ML service without blocking the core track.

The key rule:

> *Architectural work must not block progress on the sniffer and flow generator.*

---

## **Why This Approach**

I realized that if I jump straight into full microservices with events and ML, I’ll get **stuck in analysis paralysis**.
By separating concerns:

* I get **early wins** and working code to keep me motivated.
* I can **evolve the architecture gradually** instead of trying to perfect it from day one.
* I’ll have **something to show on my CV** even if the larger vision takes longer.

---

## **Architecture Notes**

Here’s the high-level structure I decided for now:

```
/docs                # Docusaurus project for documentation
/internal            # Core Go logic: sniffer, flow generator
/services
   /api              # Go backend
   /ml-api           # Python FastAPI service for ML
/web                 # React + TS dashboard
```

* **Sniffer + Flow Generator** will live in `/internal`.
* Backend and ML services will be separate but will integrate later.
* `/docs` will hold:

    * Formal documentation
    * My **journal entries** to show thought process
    * Architecture diagrams and decisions

---

## **Next Steps**

My immediate goal is to **kick off development tomorrow**:

1. Continue the work with `/internal/sniffer`.
2. Create `/docs` and write the project vision.
