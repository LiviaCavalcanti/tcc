# Trabalho de Conclusão de Curso - UFCG

Repositório com os experimentos do estudo para trabalho de conclusão de curso "Sobre oportunidades de redução de custo para o serviço AWS S3 através de alocações personalizadas". 
Para gerar amostras de teste execute `simul_tcc_py/sample_generator.py`. Nele existem várias opções para gerar diferentes tipos de amostras. Elas variam tanto no tipo da distribuição usada, quanto nos perfis abordados pelo TCC.
Esse script gera carga de trabalho para objetos no S3. Para cada dia, ele considera a quantidade de POST e GET.

Em `simul_tcc_py/`, também tem análises exploratórias feitas sobre o assunto. 

Um simulador de custo do S3 e do modelo proposto no trabalho está em `tree_algorithm.go`. Para executar a carga desejada, mude o arquivo de entrada definido no próprio script.

Obs.: Ambos os scripts não recebem argumentos e podem ser reproduzidos por
```
python simul_tcc_py/sample_generator.py
go tree_algorithm.go
```
Caso seja necessário alguma mudança, modifique o arquivo. 
