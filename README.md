# Evermos Backend Assessment

## Task 1

### What happened
During the 12.12 event (flash-sale), customers experienced order cancellations due to misreported inventory quantities data. This resulted in customer bad reviews.

### Why it happened
- **Inaccurate inventory management**
  Refers to discrepancies between the actual and recorded stock levels of products. This lead to various issues such as stockouts and inefficient order fulfillment.
- **High demand during 12.12 event**
  The 12.12 event is characterized by time-limited, heavily discounted offers, creating a sense of urgency among customers and often leading to high demand and rapid stock depletion.

### Purpose Solutions
- **Stock locking**
  Process where a certain quantity of inventory is reserved for a customer. This prevents the stock from being sold to another customer, ensuring that the order can be fulfilled and guarantee product availability for customers.
- **Purchase queue**
  Manages the order fulfillment process based on the order in which the purchases were made (First-come, First-served). This ensures that orders are processed in the sequence they were received, preventing inventory issues and order cancellations.

### Developing Guidelines
- Refer to [this document](./docs/HOW-TO-USE.md) to see project prerequisite and useful command.
- Refer to [this document](./docs/DEVELOPMENT.md) to see development decision.

### API Documentation
- Refer to this `protobuf/swagger/protobuf/api/` directory to see API Request and Response.

---

## Task 2

Run the following command:
```
go run cmd/task/main.go
```

**Sample Input**
```
5
4
#.###
..###
##.##
####.
```

**Sample Output**
```
2
```
