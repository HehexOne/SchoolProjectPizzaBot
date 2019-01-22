from django.shortcuts import render, HttpResponse, redirect
from PizzaOrders.models import Order
from django.forms import ModelForm
from django.views.decorators.csrf import csrf_exempt

# Create your views here.

rest = {1: "Дада",
        2: "Тагир",
        3: "Азяо"}
pizza = {1: "Маргарита",
         2: "Барбекю",
         3: "Сырная"}


class OrderForm(ModelForm):
    class Meta:
        model = Order
        fields = ['address', 'pizza', 'rest', 'owner_id', 'owner_platform']


@csrf_exempt
def show_db(request):
    if request.method == "POST":
        form = OrderForm(request.POST)
        if form.is_valid():
            form.save()
            lobj = Order.objects.latest('id')
            return HttpResponse(lobj.id)
    else:
        objects = Order.objects.all()
        for obj in objects:
            obj.rest = rest[obj.rest]
            obj.pizza = pizza[obj.pizza]
        return render(request, "index.html", {"orders": reversed(objects)})


def add_note(request, address, pizza, rest):
    form = OrderForm({"address": address,
                      "pizza": pizza,
                      "rest": rest})
    if form.is_valid():
        form.save()
    return show_db(request)


def delete_all(request):
    Order.objects.all().delete()
    return redirect("/")
