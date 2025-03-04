import csv
import random
from datetime import datetime, timedelta

categories = {
    "Shopping": ["Groceries", "Dinner out", "Coffee shop", "Fast food", "Supermarket"],
    "Fun": ["Movie night", "Concert ticket", "Streaming subscription", "Video game"],
    "Transport": ["Gas refill", "Uber ride", "Bus ticket", "Car maintenance"],
    "Health": ["Doctor visit", "Medicine", "Gym membership", "Dental check-up"],
    "Education": ["Online course", "Books", "School fees", "Workshop"],
    "Others": ["Home repair", "Clothing", "Electronics", "Subscription service"],
}

start_date = datetime(2024, 1, 1)
end_date = datetime(2025, 3, 4)
num_entries = 200

# Generar datos de prueba
data = []
for _ in range(num_entries):
    category = random.choice(list(categories.keys()))
    description = random.choice(categories[category])
    amount = random.randint(500, 100000)
    days_offset = random.randint(0, (end_date - start_date).days)
    date = (start_date + timedelta(days=days_offset)).strftime("%Y-%m-%d")
    data.append([description, category, amount, date])

csv_filename = "expense_data.csv"

with open(csv_filename, mode="w", newline="") as file:
    writer = csv.writer(file)
    writer.writerow(["description", "category", "amount", "date"])
    writer.writerows(data)

