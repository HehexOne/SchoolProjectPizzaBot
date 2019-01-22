# Generated by Django 2.1.5 on 2019-01-21 20:56

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('PizzaOrders', '0003_order_time'),
    ]

    operations = [
        migrations.AddField(
            model_name='order',
            name='owner_id',
            field=models.CharField(default=1, max_length=20, verbose_name='ID Владельца'),
            preserve_default=False,
        ),
        migrations.AddField(
            model_name='order',
            name='owner_platform',
            field=models.CharField(default=1, max_length=40, verbose_name='Платформа заказа'),
            preserve_default=False,
        ),
    ]