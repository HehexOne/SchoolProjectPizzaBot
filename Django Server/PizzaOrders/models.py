from django.db import models


class Order(models.Model):
    owner_platform = models.CharField(verbose_name="Платформа заказа", max_length=40)
    owner_id = models.CharField(verbose_name="ID Владельца", max_length=512)
    rest = models.IntegerField(verbose_name="ID Ресторана")
    pizza = models.IntegerField(verbose_name="ID пиццы")
    address = models.CharField(verbose_name="Адрес", max_length=512)
    time = models.DateTimeField(verbose_name="Время и дата заказа", auto_now_add=True)

    def __str__(self):
        return f"{self.owner_platform}:{self.owner_id}:{self.id}:{self.time}"
