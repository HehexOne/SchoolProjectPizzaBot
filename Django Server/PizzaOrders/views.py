from django.shortcuts import render
from PizzaOrders.models import Order
from django.forms import ModelForm
from django.views.decorators.csrf import csrf_exempt

# Create your views here.


class OrderForm(ModelForm):
    class Meta:
        model = Order
        fields = ['address', 'pizza', 'rest']


@csrf_exempt
def show_db(request):
    if request.method == "POST":
        form = OrderForm(request.POST)
        if form.is_valid():
            form.save()
    return render(request, "index.html", {"orders": Order.objects.all()})


def add_note(request, address, pizza, rest):
    form = OrderForm({"address": address, "pizza": pizza, "rest": rest})
    if form.is_valid():
        form.save()
    return show_db(request)
