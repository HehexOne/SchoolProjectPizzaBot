from django.db import models


class Order(models.Model):
    rest = models.IntegerField(verbose_name="ID Ресторана")
    pizza = models.IntegerField(verbose_name="ID пиццы")
    address = models.CharField(verbose_name="Адрес", max_length=512)
    time = models.DateTimeField(verbose_name="Время и дата заказа", auto_now_add=True)

    def __str__(self):
        return f"{self.id}:{self.time}"
